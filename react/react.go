package react

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
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

func ChainOfThought(ctx context.Context, request *structs.CompletionRequest, allowObservation bool) error {
	logger := zerolog.Ctx(ctx)

	logger.Info().Msg("detecting chain of thought")

	request.Mode = "chain_of_thought"
	req := openai.ChatCompletionRequest{
		Model:    config.OpenAI.DefaultModel.ChainOfThought,
		Messages: request.ToChatCompletionMessages(ctx, true),
	}

	var resp openai.ChatCompletionResponse

	respi, err := utils.WaitTimeout(ctx, config.Timeouts.ChainOfThought, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateChatCompletion(ctx, req)
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

	content := resp.Choices[0].Message.Content

	var action, thought, observation, answer, params string

	for ptnName := range ptns {
		if ptnName == "answer" {
			match := ptns[ptnName].FindStringSubmatch(content)
			if match == nil {
				continue
			}
			for i, name := range ptns[ptnName].SubexpNames() {
				if i != 0 && name != "" {
					switch name {
					case "answer":
						answer = match[i]
					}
				}
			}
		} else {
			for _, line := range strings.Split(content, "\n") {
				match := ptns[ptnName].FindStringSubmatch(line)
				if match == nil {
					continue
				}
				for i, name := range ptns[ptnName].SubexpNames() {
					if i != 0 && name != "" {
						switch ptnName {
						case "action":
							switch name {
							case "action":
								action = strings.ToLower(match[i])
							case "params":
								params = match[i]
							}
						case "thought":
							switch name {
							case "thought":
								thought = match[i]
							}
						case "observation":
							switch name {
							case "observation":
								observation = match[i]
							}
						}
					}
				}
			}
		}
	}

	if answer == "" && (action == "" || action == "none" || params == "" || action == "Action") {
		logger.Info().Str("action", action).Str("params", params).Str("thought", thought).Msg("chain of thought not found")
		return fmt.Errorf("no action found in response: %s", content)
	}

	if len(answer) > 0 {
		action = "answer"
		params = answer
	}

	msgs := memory.MessagesFromOpenAIChatCompletionResponse(ctx, request.User, request.Message.Interface, &resp)
	for _, msg := range msgs {
		msg.Content = params
		msg.ChainOfThought = content
		if err := msg.Save(ctx); err != nil {
			return err
		}
	}

	logger.Info().Str("action", action).
		Str("params", params).
		Str("thought", thought).
		Str("observation", observation).
		Str("answer", answer).
		Msg("chain of thought")

	return HandleAction(ctx, request, action, params, allowObservation)
}

func storeChainOfThoughtResult(ctx context.Context, request *structs.CompletionRequest, params string, content string) error {
	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.Message.Interface)
	nmsg.Role = "user"
	nmsg.Content = params
	nmsg.ChainOfThought = content
	nmsg.User = request.User

	return nmsg.Save(ctx)
}
