package react

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	reactOpenAI "github.com/kamushadenes/chloe/react/openai"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/timeouts"
	"github.com/kamushadenes/chloe/utils"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
)

var CheckpointMarker = "###CHECKPOINT###"

func DetectAction(request *structs.CompletionRequest) (*structs.ActionRequest, error) {
	logger := logging.GetLogger()

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
	j := utils.ExtractJSON(content)

	logger.Debug().Str("content", content).Str("json", j).Msg("action detection response")

	var actResp DetectedAction

	if err := json.Unmarshal([]byte(j), &actResp); err != nil {
		return nil, err
	}

	if actResp.Command.Name == "" || actResp.Command.Name == "none" {
		logger.Info().Msg("action not found")
		return nil, fmt.Errorf("no action found in response: %s", content)
	}

	msgs := memory.MessagesFromOpenAIChatCompletionResponse(request.Message.User, request.Message.Interface, &resp)
	for _, msg := range msgs {
		j = utils.ExtractJSON(msg.Content)
		msg.SetContent(j)
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

	logger.Info().
		EmbedObject(actResp).
		Msg("action detected")

	if config.React.ReportThoughts {
		_ = request.Message.SendText(fmt.Sprintf("*Chain of Thought:\n  - %s*", strings.Join(actResp.Thoughts.ChainOfThought, "\n  - ")), false)
		_ = request.Message.SendText(fmt.Sprintf("*Plan:\n  - %s*", strings.Join(actResp.Thoughts.Plan, "\n  - ")), false)
		_ = request.Message.SendText(fmt.Sprintf("*Criticism: %s*", actResp.Thoughts.Criticism), false)
	}

	actReq := structs.NewActionRequest()
	actReq.ID = request.ID
	actReq.Message = request.Message
	actReq.Context = request.Context
	actReq.Action = actResp.Command.Name
	actReq.Params = actResp.Command.Params
	actReq.Thought = strings.Join(actResp.Thoughts.ChainOfThought, "\n - ")
	actReq.Writers = []io.WriteCloser{request.Writer}

	logger.Info().
		EmbedObject(actResp).
		Float64("estimatedPromptCost",
			config.OpenAI.GetModel(config.Completion).GetChatCompletionCost(req.Messages, "")).
		Float64("estimatedResponseCost",
			config.OpenAI.GetModel(config.Completion).GetChatCompletionCost(nil, content)).
		Msg("action detection finished")

	return actReq, nil
}
