package react

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/interfaces/telegram"
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

func HandleAction(ctx context.Context, request *structs.CompletionRequest, action string, params string, allowObservation bool) error {
	tokenCount := request.CountTokens(request.ToChatCompletionMessages(ctx, true))
	truncateTokenCount := config.OpenAI.MaxTokens[config.OpenAI.DefaultModel.ChainOfThought] - tokenCount
	if truncateTokenCount < config.OpenAI.MinReplyTokens[config.OpenAI.DefaultModel.ChainOfThought] {
		truncateTokenCount = config.OpenAI.MinReplyTokens[config.OpenAI.DefaultModel.ChainOfThought]
	}

	switch strings.ToLower(action) {
	case "observation", "answer":
		if !allowObservation {
			return NotifyError(request, fmt.Errorf("observation is not allowed yet"))
		}
		_, err := request.Writer.Write([]byte(params))

		return err
	case "google":
		var bw = BytesWriter{}

		req := structs.NewScrapeRequest()
		req.Context = ctx
		req.Writer = &bw
		req.SkipClose = false
		req.User = request.User
		req.Content = params

		if err := Google(ctx, req); err != nil {
			return err
		}

		content := fmt.Sprintf("Observation: %s", Truncate(string(bw.Bytes), truncateTokenCount))

		if err := storeChainOfThoughtResult(ctx, request, params, content); err != nil {
			return err
		}

		var nreq = request.Copy()
		nreq.Message.Content = fmt.Sprintf(content)

		return ChainOfThought(ctx, nreq, true)
	case "scrape", "web":
		var bw = BytesWriter{}

		req := structs.NewScrapeRequest()
		req.Context = ctx
		req.Writer = &bw
		req.SkipClose = false
		req.User = request.User
		req.Content = params

		if err := Scrape(ctx, req); err != nil {
			return err
		}

		content := fmt.Sprintf("Observation: %s", Truncate(string(bw.Bytes), truncateTokenCount))

		if err := storeChainOfThoughtResult(ctx, request, params, content); err != nil {
			return err
		}

		var nreq = request.Copy()
		nreq.Message.Content = fmt.Sprintf("Observation: %s", content)

		return ChainOfThought(ctx, nreq, true)
	case "image", "dalle", "dall-e":
		var ws []io.WriteCloser

		if request.Message.Interface == "telegram" {
			iw := request.Writer.(*telegram.TelegramWriter).ToImageWriter()
			for k := 0; k < 4; k++ {
				ws = append(ws, iw.(*telegram.TelegramWriter).Subwriter())
			}
		} else {
			ws = append(ws, request.Writer)
		}

		errorCh := make(chan error)
		req := structs.NewGenerationRequest()
		req.Context = ctx
		req.User = request.User
		req.Prompt = params
		req.StartChannel = request.StartChannel
		req.ContinueChannel = request.ContinueChannel
		req.ErrorChannel = errorCh

		req.Writers = ws

		channels.GenerationRequestsCh <- req

		return <-errorCh
	case "audio", "tts", "speak":
		var ws []io.WriteCloser

		if request.Message.Interface == "telegram" {
			iw := request.Writer.(*telegram.TelegramWriter).ToAudioWriter()
			ws = append(ws, iw)
		} else {
			ws = append(ws, request.Writer)
		}

		errorCh := make(chan error)

		req := structs.NewTTSRequest()
		req.Context = ctx
		req.User = request.User
		req.Content = params
		req.ErrorChannel = errorCh

		req.Writers = ws

		channels.TTSRequestsCh <- req

		return <-errorCh
	case "calculate", "math":
		var bw = BytesWriter{}

		req := structs.NewCalculationRequest()
		req.Context = ctx
		req.Writer = &bw
		req.User = request.User
		req.Content = params

		if err := Calculate(ctx, req); err != nil {
			return NotifyError(request, err)
		}

		content := fmt.Sprintf("Observation: %s", Truncate(string(bw.Bytes), truncateTokenCount))

		if err := storeChainOfThoughtResult(ctx, request, params, content); err != nil {
			return err
		}

		var nreq = request.Copy()
		nreq.Message.Content = fmt.Sprintf("Observation: %s", content)

		return ChainOfThought(ctx, nreq, true)
	default:
		return fmt.Errorf("unknown action %s", action)
	}
}
