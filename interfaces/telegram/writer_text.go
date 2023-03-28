package telegram

import (
	"bytes"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
)

func (w *TelegramWriter) ToTextWriter() *TelegramWriter {
	_, _ = w.Bot.Send(tgbotapi.NewChatAction(w.ChatID, tgbotapi.ChatTyping))
	return NewTextWriter(w.Context, w.Request, w.ReplyID != 0)
}

func NewTextWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *TelegramWriter {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Telegram.API,
		ChatID:  request.GetMessage().Source.Telegram.Update.Message.Chat.ID,
		Type:    "text",
		Request: request,
		bufs:    []bytes.Buffer{{}},
		bufID:   0,
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Telegram.Update.Message.MessageID
	}
	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	return w
}

func (w *TelegramWriter) closeText() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	logger.Debug().Int64("chatID", w.ChatID).Msg("replying with text")

	msgs := utils.StringToChunks(w.bufs[0].String(), config.Discord.MaxMessageLength)

	if w.externalID == 0 {
		for k := range msgs {
			if err := w.Request.GetMessage().SendText(w.bufs[k].String(), true, w.ReplyID); err != nil {
				return err
			}
		}
	} else {
		if _, err := w.Bot.Send(tgbotapi.NewEditMessageText(w.ChatID, w.externalID, msgs[0])); err != nil {
			return err
		}

		for k := range msgs[1:] {
			if err := w.Request.GetMessage().SendText(msgs[k], true); err != nil {
				return err
			}
		}
	}

	return nil
}
