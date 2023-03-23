package openai

import (
	"context"
	"errors"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"strings"
)

func Complete(ctx context.Context, request *structs.CompletionRequest) error {
	logger := zerolog.Ctx(ctx)
	ctx = logger.WithContext(ctx)

	ocontent := request.Message.Content

	cotreq := request.Copy()
	cotreq.Message.Content = fmt.Sprintf("Question: %s", ocontent)
	err := react.ChainOfThought(ctx, cotreq, false)
	if err == nil {
		if !request.SkipClose {
			err := request.Writer.Close()
			return react.NotifyError(request, err)
		}
		return react.NotifyError(request, nil)
	}

	req := openai.ChatCompletionRequest{
		Model:    config.OpenAI.DefaultModel.Completion,
		Messages: request.ToChatCompletionMessages(ctx, false),
	}

	logger.Info().Int("messagesInContext", len(req.Messages)).Msg("requesting completion")

	var stream *openai.ChatCompletionStream

	respi, err := utils.WaitTimeout(ctx, config.Timeouts.Completion, func(ch chan interface{}, errCh chan error) {
		stream, err := openAIClient.CreateChatCompletionStream(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error requesting completion")
			errCh <- err
		}
		ch <- stream
	})
	if err != nil {
		return react.NotifyError(request, err)
	}

	stream = respi.(*openai.ChatCompletionStream)

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

	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
	nmsg.Content = responseMessage
	nmsg.Role = "assistant"
	nmsg.User = request.User

	_ = nmsg.Save(ctx)

	react.WriteResult(request, responseMessage)

	if !request.SkipClose {
		err := request.Writer.Close()
		return react.NotifyError(request, err)
	}

	return react.NotifyError(request, nil)
}
