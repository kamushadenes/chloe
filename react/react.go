package react

import (
	"encoding/json"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"regexp"
	"strings"
)

var openAIClient = openai.NewClient(config.OpenAI.APIKey)

var ptns = map[string]*regexp.Regexp{
	"action":      regexp.MustCompile(`^[Aa]ction:\s*(?P<action>\w+): (?P<params>.*)$`),
	"thought":     regexp.MustCompile(`^[Tt]hought:\s*(?P<thought>.*)$`),
	"observation": regexp.MustCompile(`^[Oo]bservation:\s*(?P<observation>.*)$`),
	"answer":      regexp.MustCompile(`^[Aa]nswer:\s*(?P<answer>.*)`),
}

func ChainOfThought(request *structs.CompletionRequest) error {
	logger := zerolog.Ctx(request.Context)

	logger.Info().Msg("detecting chain of thought")

	request.Mode = "chain_of_thought"
	req := openai.ChatCompletionRequest{
		Model:    config.OpenAI.DefaultModel.ChainOfThought,
		Messages: request.ToChatCompletionMessages(),
	}

	var resp openai.ChatCompletionResponse

	respi, err := utils.WaitTimeout(request.Context, config.Timeouts.ChainOfThought, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateChatCompletion(request.Context, req)
		if err != nil {
			logger.Error().Err(err).Msg("error requesting chain of thought")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return err
	}

	resp = respi.(openai.ChatCompletionResponse)

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	var cotResp ChainOfThoughtResponse
	if err := json.Unmarshal([]byte(content), &cotResp); err != nil {
		return err
	}

	if cotResp.Action == "" || cotResp.Action == "none" {
		logger.Info().Msg("chain of thought not found")
		return fmt.Errorf("no action found in response: %s", content)
	}

	msgs := memory.MessagesFromOpenAIChatCompletionResponse(request.Context, request.Message.User, request.Message.Interface, &resp)
	for _, msg := range msgs {
		msg.Content = content // params
		// msg.ChainOfThought = content
		if err := msg.Save(request.Context); err != nil {
			return err
		}
	}

	logger.Info().Str("action", cotResp.Action).
		Str("params", cotResp.Params).
		Str("thought", cotResp.Thought).
		Msg("chain of thought")

	actReq := structs.NewActionRequest()
	actReq.ID = request.ID
	actReq.Message = request.Message
	actReq.Context = request.Context
	actReq.Action = cotResp.Action
	actReq.Params = cotResp.Params
	actReq.Thought = cotResp.Thought
	actReq.Writers = []io.WriteCloser{request.Writer}

	return HandleAction(actReq)
}

func storeChainOfThoughtResult(request structs.ActionOrCompletionRequest, content string) error {
	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.GetMessage().Interface)
	nmsg.Role = "assistant"
	nmsg.Content = content
	nmsg.User = request.GetMessage().User

	return nmsg.Save(request.GetContext())
}
