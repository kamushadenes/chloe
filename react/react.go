package react

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	reactOpenAI "github.com/kamushadenes/chloe/react/openai"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/timeouts"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
)

var CheckpointMarker = "###CHECKPOINT###"

func DetectAction(request *structs.CompletionRequest) (*structs.ActionRequest, error) {
	logger := zerolog.Ctx(request.Context)

	logger.Info().Msg("detecting action")

	request.Mode = "action_detection"
	req := openai.ChatCompletionRequest{
		Model:    config.OpenAI.DefaultModel.ChainOfThought.String(),
		Messages: request.ToChatCompletionMessages(),
	}

	var resp openai.ChatCompletionResponse

	respi, err := timeouts.WaitTimeout(request.Context, config.Timeouts.ChainOfThought, func(ch chan interface{}, errCh chan error) {
		resp, err := reactOpenAI.OpenAIClient.CreateChatCompletion(request.Context, req)
		if err != nil {
			logger.Error().Err(err).Msg("error detecting action")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return nil, err
	}

	resp = respi.(openai.ChatCompletionResponse)

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	var cotResp DetectedAction
	if err := json.Unmarshal([]byte(content), &cotResp); err != nil {
		return nil, err
	}

	if cotResp.Action == "" || cotResp.Action == "none" {
		logger.Info().Msg("action not found")
		return nil, fmt.Errorf("no action found in response: %s", content)
	}

	msgs := memory.MessagesFromOpenAIChatCompletionResponse(request.Message.User, request.Message.Interface, &resp)
	for _, msg := range msgs {
		msg.SetContent(fmt.Sprintf("Thought: %s\nAction: %s\nParams: %s", cotResp.Thought, cotResp.Action, cotResp.Params))
		if err := msg.Save(request.Context); err != nil {
			return nil, err
		}
	}

	msg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "internal")
	msg.User = request.Message.User
	msg.Context = request.Message.Context
	msg.SetContent(CheckpointMarker)
	msg.Role = "user"
	if err := msg.Save(request.Context); err != nil {
		return nil, err
	}

	logger.Info().Str("action", cotResp.Action).
		Str("params", cotResp.Params).
		Str("thought", cotResp.Thought).
		Msg("action detected")

	actReq := structs.NewActionRequest()
	actReq.ID = request.ID
	actReq.Message = request.Message
	actReq.Context = request.Context
	actReq.Action = cotResp.Action
	actReq.Params = cotResp.Params
	actReq.Thought = cotResp.Thought
	actReq.Writers = []io.WriteCloser{request.Writer}

	logger.Info().
		Str("action", cotResp.Action).
		Str("params", cotResp.Params).
		Str("thought", cotResp.Thought).
		Float64("estimatedPromptCost",
			config.OpenAI.GetModel(config.Completion).GetChatCompletionCost(req.Messages, "")).
		Float64("estimatedResponseCost",
			config.OpenAI.GetModel(config.Completion).GetChatCompletionCost(nil, content)).
		Msg("action dectection finished")

	return actReq, nil
}
