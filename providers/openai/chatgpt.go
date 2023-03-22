package openai

import (
	"context"
	"errors"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"strings"
)

func Complete(ctx context.Context, request *structs.CompletionRequest) error {
	logger := zerolog.Ctx(ctx)
	ctx = logger.WithContext(ctx)

	cotreq := request.Copy()
	cotreq.Content = fmt.Sprintf("Question: %s", cotreq.Content)
	err := react.ChainOfThought(ctx, cotreq, false)
	if err == nil {
		if !request.SkipClose {
			err := request.Writer.Close()
			return react.NotifyError(request, err)
		}
		return react.NotifyError(request, nil)
	}

	req := openai.ChatCompletionRequest{
		Model:    config.OpenAI.DefaultModel[config.ModelPurposeCompletion],
		Messages: request.ToChatCompletionMessages(ctx, false),
	}

	logger.Info().Int("messagesInContext", len(req.Messages)).Msg("requesting completion")

	stream, err := openAIClient.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return react.NotifyError(request, err)
	}

	react.StartAndWait(request)

	writeStatusCode(request.Writer, http.StatusOK)

	var responseMessage string

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return react.NotifyError(request, err)
		}

		content := response.Choices[0].Delta.Content

		responseMessage += content
		_, _ = request.Writer.Write([]byte(content))

		flush(request.Writer)
	}

	responseMessage = strings.TrimSpace(responseMessage)
	_ = memory.SaveMessage(ctx, request.User.ID, "assistant", responseMessage, "")

	react.WriteResult(request, responseMessage)

	if !request.SkipClose {
		err := request.Writer.Close()
		return react.NotifyError(request, err)
	}

	return react.NotifyError(request, nil)
}
