package openai

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/media"
	putils "github.com/kamushadenes/chloe/providers/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/timeouts"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

// newTranscriptionRequest creates a new openai.AudioRequest for transcription.
func newTranscriptionRequest(request *structs.TranscriptionRequest) openai.AudioRequest {
	return openai.AudioRequest{
		Model:    config.OpenAI.DefaultModel.Transcription.String(),
		FilePath: request.FilePath,
	}
}

// createTranscriptionRequestWithTimeout attempts to create an AudioResponse with a timeout.
// Returns the created AudioResponse or an error if the request times out or encounters an issue.
func createTranscriptionRequestWithTimeout(request *structs.TranscriptionRequest, req openai.AudioRequest) (openai.AudioResponse, error) {
	logger := logging.GetLogger().With().Str("file", request.FilePath).Logger()

	respi, err := timeouts.WaitTimeout(request.GetContext(), config.Timeouts.Transcription, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateTranscription(request.GetContext(), req)
		if err != nil {
			logger.Error().Err(err).Msg("error transcribing audio")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return openai.AudioResponse{}, err
	}

	return respi.(openai.AudioResponse), err
}

// processSuccessfulTranscriptionRequest processes a successful transcription request, writing the response
// to the given request.Writer and, if present, sending the result to request.ResultChannel.
// Returns an error if there's an issue during the process.
func processSuccessfulTranscriptionRequest(request *structs.TranscriptionRequest, response openai.AudioResponse) error {
	channels.StartAndWait(request)

	putils.WriteStatusCode(request.Writer, http.StatusOK)

	_, err := request.Writer.Write([]byte(response.Text))
	if err != nil {
		return err
	}

	if request.ResultChannel != nil {
		request.ResultChannel <- fmt.Sprintf("Transcription: %s", response.Text)
	}

	return nil
}

// recordTranscriptionResponse records the transcription response as a new message in the memory.
// Returns an error if there's an issue during the process.
func recordTranscriptionResponse(request *structs.TranscriptionRequest, response openai.AudioResponse) error {
	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
	nmsg.Role = "user"
	nmsg.User = request.GetMessage().User
	nmsg.SetContent(response.Text)

	return nmsg.Save(request.GetContext())
}

// Transcribe processes a transcription request for an audio file using the OpenAI API.
// Returns an error if there's an issue during the process.
func Transcribe(request *structs.TranscriptionRequest) error {
	logger := structs.LoggerFromRequest(request)

	if duration, err := media.GetMediaDuration(request.FilePath); err != nil {
		logger.Error().Err(err).Msg("error getting media duration, skipping cost calculation")
		logger.Info().Msg("transcribing file")
	} else {
		price := config.OpenAI.GetModel(config.Transcription).UsageCost.Price
		totalCost := price * duration.Minutes()

		cost.AddCategoryCost(string(config.Transcription), totalCost)

		logger.Info().Float64("cost", totalCost).Dur("duration", duration).Msg("transcribing file")
	}

	req := newTranscriptionRequest(request)

	response, err := createTranscriptionRequestWithTimeout(request, req)
	if err != nil {
		return channels.NotifyError(request, errors.ErrTranscriptionFailed, err)
	}

	err = processSuccessfulTranscriptionRequest(request, response)
	if err != nil {
		return channels.NotifyError(request, errors.ErrTranscriptionFailed, err)
	}

	err = recordTranscriptionResponse(request, response)
	if err != nil {
		err = errors.Wrap(errors.ErrTranscriptionFailed, err)
	}

	return channels.NotifyError(request, err)
}
