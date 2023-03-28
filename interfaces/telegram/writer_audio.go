package telegram

import (
	"bytes"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
)

func (w *TelegramWriter) ToAudioWriter() *TelegramWriter {
	_, _ = w.Bot.Send(tgbotapi.NewChatAction(w.ChatID, tgbotapi.ChatRecordVoice))
	return NewAudioWriter(w.Context, w.Request, w.ReplyID != 0)
}

func NewAudioWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *TelegramWriter {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Telegram.API,
		ChatID:  request.GetMessage().Source.Telegram.Update.Message.Chat.ID,
		Type:    "audio",
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

func (w *TelegramWriter) closeAudio() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	logger.Debug().Int64("chatID", w.ChatID).Msg("replying with audio")
	tmsg := tgbotapi.NewVoice(w.ChatID, tgbotapi.FileReader{
		Reader: bytes.NewReader(w.bufs[0].Bytes()),
	})
	tmsg.ReplyToMessageID = w.ReplyID
	_, err := w.Bot.Send(tmsg)
	return err
}
