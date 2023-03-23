package openai

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func Transcribe(ctx context.Context, request *structs.TranscriptionRequest) error {
	logger := zerolog.Ctx(ctx).With().Str("file", request.FilePath).Logger()

	ctx = logger.WithContext(ctx)

	logger.Info().Msg("transcribing file")

	req := openai.AudioRequest{
		Model:    config.OpenAI.DefaultModel.Transcription,
		FilePath: request.FilePath,
	}

	var response openai.AudioResponse

	respi, err := utils.WaitTimeout(ctx, config.Timeouts.Transcription, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateTranscription(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error transcribing audio")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return react.NotifyError(request, err)
	}

	response = respi.(openai.AudioResponse)

	react.StartAndWait(request)

	writeStatusCode(request.Writer, http.StatusOK)

	_, err = request.Writer.Write([]byte(response.Text))
	defer request.Writer.Close()

	if request.ResultChannel != nil {
		request.ResultChannel <- fmt.Sprintf("Transcription: %s", response.Text)
	}

	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
	nmsg.Role = "user"
	nmsg.User = request.User
	nmsg.Content = response.Text

	return react.NotifyError(request, nmsg.Save(ctx))
}
