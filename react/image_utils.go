package react

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/interfaces/discord"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/resources"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
)

func imagePreActions(a Action, request *structs.ActionRequest) error {
	var ws []io.WriteCloser

	switch request.Message.Interface {
	case "telegram":
		w := request.GetWriters()[0].(*telegram.TelegramWriter)
		iw := w.ToImageWriter()
		for k := 0; k < config.Telegram.ImageCount; k++ {
			siw := iw.Subwriter()
			siw.SetPrompt(request.Params)
			ws = append(ws, siw)
		}
	case "discord":
		w := request.GetWriters()[0].(*discord.DiscordWriter)
		iw := w.ToImageWriter()
		for k := 0; k < config.Discord.ImageCount; k++ {
			siw := iw.Subwriter()
			siw.SetPrompt(request.Params)
			ws = append(ws, siw)
		}
	default:
		ws = append(ws, request.GetWriters()[0])
	}

	a.SetWriters(ws)

	if config.React.ImproveImagePrompts {
		b := &BytesWriter{}

		na := NewMidjourneyPromptGeneratorAction()
		na.SetParams(a.GetParams())
		na.SetWriters([]io.WriteCloser{b})
		request.Message.NotifyAction(na.GetNotification())
		if err := na.Execute(request); err == nil {
			a.SetParams(string(b.Bytes))
		}
	}

	return nil
}

func improvePrompt(request *structs.ActionRequest) string {
	logger := logging.GetLogger()

	prompt, err := resources.GetPrompt("midjourney_prompt_generator", &resources.PromptArgs{
		Args: map[string]interface{}{},
		Mode: "midjourney_prompt_generator",
	})

	req := openai.ChatCompletionRequest{
		Model: config.OpenAI.DefaultModel.ChainOfThought,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: request.Params,
			},
		},
	}

	var resp openai.ChatCompletionResponse

	respi, err := utils.WaitTimeout(request.Context, config.Timeouts.ChainOfThought, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateChatCompletion(request.Context, req)
		if err != nil {
			logger.Error().Err(err).Msg("error requesting prompt improvement")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return request.Params
	}

	resp = respi.(openai.ChatCompletionResponse)

	content := strings.TrimSpace(resp.Choices[0].Message.Content)

	return content
}
