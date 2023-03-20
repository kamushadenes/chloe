package openai

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"net/http"
)

func Transcribe(ctx context.Context, request *structs.TranscriptionRequest) error {
	logger := zerolog.Ctx(ctx).With().Str("file", request.FilePath).Logger()

	ctx = logger.WithContext(ctx)

	logger.Info().Msg("transcribing file")

	req := openai.AudioRequest{
		Model:    config.OpenAI.DefaultModel[config.ModelPurposeTranscription],
		FilePath: request.FilePath,
	}

	response, err := openAIClient.CreateTranscription(ctx, req)
	if err != nil {
		return react.NotifyError(request, err)
	}

	react.StartAndWait(request)

	writeStatusCode(request.Writer, http.StatusOK)

	_, err = request.Writer.Write([]byte(response.Text))
	defer request.Writer.Close()

	if request.ResultChannel != nil {
		request.ResultChannel <- response.Text
	}

	err = memory.SaveMessage(ctx, request.User.ID, "user", response.Text, "")

	return react.NotifyError(request, err)
}
