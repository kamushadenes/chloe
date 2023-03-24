package openai

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	putils "github.com/kamushadenes/chloe/providers/utils"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

// newTranscriptionRequest creates a new openai.AudioRequest for transcription.
func newTranscriptionRequest(ctx context.Context, request *structs.TranscriptionRequest) openai.AudioRequest {
	return openai.AudioRequest{
		Model:    config.OpenAI.DefaultModel.Transcription,
		FilePath: request.FilePath,
	}
}

// createTranscriptionRequestWithTimeout attempts to create an AudioResponse with a timeout.
// Returns the created AudioResponse or an error if the request times out or encounters an issue.
func createTranscriptionRequestWithTimeout(ctx context.Context, request *structs.TranscriptionRequest, req openai.AudioRequest) (openai.AudioResponse, error) {
	logger := zerolog.Ctx(ctx).With().Str("file", request.FilePath).Logger()
	ctx = logger.WithContext(ctx)

	respi, err := utils.WaitTimeout(ctx, config.Timeouts.Transcription, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateTranscription(ctx, req)
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
func processSuccessfulTranscriptionRequest(ctx context.Context, request *structs.TranscriptionRequest, response openai.AudioResponse) error {
	react.StartAndWait(request)

	putils.WriteStatusCode(request.Writer, http.StatusOK)

	_, err := request.Writer.Write([]byte(response.Text))
	if err != nil {
		return err
	}

	if request.ResultChannel != nil {
		request.ResultChannel <- fmt.Sprintf("Transcription: %s", response.Text)
	}

	return request.Writer.Close()
}

// recordTranscriptionResponse records the transcription response as a new message in the memory.
// Returns an error if there's an issue during the process.
func recordTranscriptionResponse(ctx context.Context, request *structs.TranscriptionRequest, response openai.AudioResponse) error {
	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
	nmsg.Role = "user"
	nmsg.User = request.User
	nmsg.Content = response.Text

	return nmsg.Save(ctx)
}

// Transcribe processes a transcription request for an audio file using the OpenAI API.
// Returns an error if there's an issue during the process.
func Transcribe(ctx context.Context, request *structs.TranscriptionRequest) error {
	logger := zerolog.Ctx(ctx).With().Str("file", request.FilePath).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("transcribing file")

	req := newTranscriptionRequest(ctx, request)

	response, err := createTranscriptionRequestWithTimeout(ctx, request, req)
	if err != nil {
		return react.NotifyError(request, err)
	}

	err = processSuccessfulTranscriptionRequest(ctx, request, response)
	if err != nil {
		return react.NotifyError(request, err)
	}

	return react.NotifyError(request, recordTranscriptionResponse(ctx, request, response))
}
