package openai

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	putils "github.com/kamushadenes/chloe/providers/utils"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/react/actions"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/timeouts"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
)

// detectAction processes the chain of thought for a completion request
// and returns an error if there is any issue.
func detectAction(request *structs.CompletionRequest) error {
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
	cotreq.Message.SetContent(string(b))

	actReq, err := react.DetectAction(cotreq)
	if err != nil {
		return err
	}

	return actions.HandleAction(actReq)
}

// newChatCompletionRequest creates a new OpenAI ChatCompletionRequest
// from the provided CompletionRequest.
func newChatCompletionRequest(request *structs.CompletionRequest) openai.ChatCompletionRequest {
	return openai.ChatCompletionRequest{
		Model:    config.OpenAI.DefaultModel.Completion.String(),
		Messages: request.ToChatCompletionMessages(),
	}
}

// createChatCompletionWithTimeout attempts to create a ChatCompletionStream with a timeout.
// Returns the created ChatCompletionStream or an error if the request times out or encounters an issue.
func createChatCompletionWithTimeout(ctx context.Context, req openai.ChatCompletionRequest) (*openai.ChatCompletionStream, error) {
	logger := logging.GetLogger()

	logger.Info().
		Float64("estimatedPromptCost",
			config.OpenAI.GetModel(config.Completion).GetChatCompletionCost(req.Messages, "")).
		Msg("creating chat completion stream")

	respi, err := timeouts.WaitTimeout(ctx, config.Timeouts.Completion, func(ch chan interface{}, errCh chan error) {
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

// processCompletionStream processes a ChatCompletionStream and writes the response to the request.Writer.
// Returns the response message as a string, or an error if there's an issue while processing the stream.
func processCompletionStream(request *structs.CompletionRequest, stream *openai.ChatCompletionStream) (string, error) {
	channels.StartAndWait(request)

	resp := stream.GetResponse()
	defer resp.Body.Close()

	putils.WriteStatusCode(request.Writer, resp.StatusCode)

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

	channels.WriteResult(request, responseMessage)

	return strings.TrimSpace(responseMessage), nil
}

// recordAssistantResponse saves the assistant's response as a message in the memory.
// Returns an error if there's an issue while saving the message.
func recordAssistantResponse(request *structs.CompletionRequest, responseMessage string) error {
	logger := logging.GetLogger()

	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
	nmsg.SetContent(responseMessage)
	nmsg.Role = "assistant"
	nmsg.User = request.GetMessage().User

	logger.Info().
		Float64("estimatedResponseCost",
			config.OpenAI.GetModel(config.Completion).GetChatCompletionCost(nil, "")).
		Msg("recording assistant response")

	return nmsg.Save(request.GetContext())
}

// Complete processes a completion request by interacting with the OpenAI API.
// Returns an error if there's an issue during the process.
func Complete(r *structs.CompletionRequest, skipCoT ...bool) error {
	request := r.Copy()
	logger := structs.LoggerFromRequest(request)

	if len(skipCoT) == 0 || !skipCoT[0] {
		if err := detectAction(request); err == nil {
			return channels.NotifyError(request, err)
		} else if errors.Is(err, errors.ErrProceed) {
			msg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
			msg.SetContent(fmt.Sprintf("summarize everything since \"%s\", never mention the checkpoint, try to make it look like a news article or a Wikipedia page, don't provide explanation", react.CheckpointMarker))
			msg.ErrorCh = request.Message.ErrorCh
			msg.Role = "user"
			msg.User = request.Message.User
			msg.Context = request.Message.Context
			msg.Source = request.Message.Source
			if err := msg.Save(request.GetContext()); err != nil {
				return errors.Wrap(errors.ErrCompletionFailed, err)
			}
			return Complete(request, true)
		} else if errors.Is(err, errors.ErrActionFailed) {
			return channels.NotifyError(request, err)
		}
	}

	req := newChatCompletionRequest(request)

	logger.Info().Int("messagesInContext", len(req.Messages)).Msg("requesting completion")

	stream, err := createChatCompletionWithTimeout(request.GetContext(), req)
	if err != nil {
		return channels.NotifyError(request, errors.ErrCompletionFailed, err)
	}

	responseMessage, err := processCompletionStream(request, stream)
	if err != nil {
		return channels.NotifyAndClose(request, request.Writer, errors.ErrCompletionFailed, err)
	}
	if responseMessage == "" {
		return channels.NotifyAndClose(request, request.Writer, errors.ErrCompletionFailed, fmt.Errorf("empty response"))
	}
	if strings.TrimSpace(responseMessage) == "" {
		_ = request.Message.User.DeleteOldestMessage(request.GetContext())
		return Complete(request, false)
	}

	err = recordAssistantResponse(request, responseMessage)
	if err != nil {
		err = errors.Wrap(errors.ErrCompletionFailed, err)
	}

	return channels.NotifyAndClose(request, request.Writer, err)
}
