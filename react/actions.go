package react

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/interfaces/discord"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"strings"
)

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
		_ = request.GetMessage().SendText(i18n.GetImageGenerationText())

		var ws []io.WriteCloser

		if request.Message.Interface == "telegram" {
			w := request.Writer.(*telegram.TelegramWriter)
			iw := w.ToImageWriter().(*telegram.TelegramWriter)
			for k := 0; k < config.Telegram.ImageCount; k++ {
				siw := iw.Subwriter()
				siw.SetPrompt(params)
				ws = append(ws, siw)
			}
		} else if request.Message.Interface == "discord" {
			w := request.Writer.(*discord.DiscordWriter)
			iw := w.ToImageWriter().(*discord.DiscordWriter)
			for k := 0; k < config.Discord.ImageCount; k++ {
				siw := iw.Subwriter()
				siw.SetPrompt(params)
				ws = append(ws, siw)
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
		} else if request.Message.Interface == "discord" {
			iw := request.Writer.(*discord.DiscordWriter)
			iw.Prompt = params
			ws = append(ws, iw.ToAudioWriter())
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
