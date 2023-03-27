package openai

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/interfaces/discord"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/memory"
	putils "github.com/kamushadenes/chloe/providers/utils"
	"github.com/kamushadenes/chloe/react"
	errors2 "github.com/kamushadenes/chloe/react/errors"
	utils2 "github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
	"strings"
)

// processChainOfThought processes the chain of thought for a completion request
// and returns an error if there is any issue.
func processChainOfThought(request *structs.CompletionRequest) error {
	ocontent := request.Message.Content

	params := struct {
		Question string `json:"question"`
	}{
		Question: ocontent,
	}

	b, err := json.Marshal(params)
	if err != nil {
		return err
	}

	cotreq := request.Copy()
	cotreq.Message.Content = string(b)
	return react.ChainOfThought(cotreq)
}

// newChatCompletionRequest creates a new OpenAI ChatCompletionRequest
// from the provided CompletionRequest.
func newChatCompletionRequest(request *structs.CompletionRequest) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:    config.OpenAI.DefaultModel.Completion,
		Messages: request.ToChatCompletionMessages(),
	}
}

// createChatCompletionWithTimeout attempts to create a ChatCompletionStream with a timeout.
// Returns the created ChatCompletionStream or an error if the request times out or encounters an issue.
func createChatCompletionWithTimeout(ctx context.Context, req openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := utils.WaitTimeout(ctx, config.Timeouts.Completion, func(ch chan interface{}, errCh chan error) {
		stream, err := openAIClient.CreateChatCompletionStream(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error requesting completion")
			errCh <- err
		}
		ch <- stream
	})
	if err != nil {
		return nil, err
	}

	return respi.(*openai.ChatCompletionStream), err
}

// processSuccessfulCompletionStream processes a ChatCompletionStream and writes the response to the request.Writer.
// Returns the response message as a string, or an error if there's an issue while processing the stream.
func processSuccessfulCompletionStream(request *structs.CompletionRequest, stream *openai.ChatCompletionStream) (string, error) {
	utils2.StartAndWait(request)

	putils.WriteStatusCode(request.Writer, http.StatusOK)

	var responseMessage string

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return "", err
		}

		content := response.Choices[0].Delta.Content

		responseMessage += content
		_, err = request.Writer.Write([]byte(content))
		if err != nil {
			return "", err
		}

		putils.Flush(request.Writer)
	}

	utils2.WriteResult(request, responseMessage)

	return strings.TrimSpace(responseMessage), nil
}

// recordAssistantResponse saves the assistant's response as a message in the memory.
// Returns an error if there's an issue while saving the message.
func recordAssistantResponse(request *structs.CompletionRequest, responseMessage string) error {
	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
	nmsg.Content = responseMessage
	nmsg.Role = "assistant"
	nmsg.User = request.GetMessage().User

	return nmsg.Save(request.GetContext())
}

// Complete processes a completion request by interacting with the OpenAI API.
// Returns an error if there's an issue during the process.
func Complete(r *structs.CompletionRequest, skipCoT ...bool) error {
	request := r.Copy()
	logger := structs.LoggerFromRequest(request)

	switch w := request.Writer.(type) {
	case *telegram.TelegramWriter:
		request.Writer = w.ToTextWriter()
	case *discord.DiscordWriter:
		request.Writer = w.ToTextWriter()
	}

	// TODO: call CoT outside of this function
	if len(skipCoT) == 0 || !skipCoT[0] {
		if err := processChainOfThought(request); err == nil {
			return utils2.NotifyError(request, err)
		} else if errors.Is(err, errors2.ErrProceed) {
			msg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
			msg.Content = "great work, with this new information, provide me an explanation of my last question in a Wikipedia page style with whatever information you have available, don't worry if you don't have enough information, I'll ask you for more"
			msg.ErrorCh = request.Message.ErrorCh
			msg.Role = "user"
			msg.User = request.Message.User
			msg.Context = request.Message.Context
			msg.Source = request.Message.Source
			if err := msg.Save(request.GetContext()); err != nil {
				return err
			}
			return Complete(request, true)
		}
	}

	req := newChatCompletionRequest(request)

	logger.Info().Int("messagesInContext", len(req.Messages)).Msg("requesting completion")

	stream, err := createChatCompletionWithTimeout(request.GetContext(), req)
	if err != nil {
		return utils2.NotifyError(request, err)
	}

	responseMessage, err := processSuccessfulCompletionStream(request, stream)
	if err != nil {
		return utils2.NotifyAndClose(request, request.Writer, err)
	}
	if strings.TrimSpace(responseMessage) == "" {
		_ = request.Message.User.DeleteOldestMessage(request.GetContext())
		return Complete(request, false)
	}

	err = recordAssistantResponse(request, responseMessage)

	return utils2.NotifyAndClose(request, request.Writer, err)
}
