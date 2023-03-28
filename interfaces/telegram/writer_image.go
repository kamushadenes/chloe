package telegram

import (
	"bytes"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
)

func (w *TelegramWriter) ToImageWriter() *TelegramWriter {
	_, _ = w.Bot.Send(tgbotapi.NewChatAction(w.ChatID, tgbotapi.ChatUploadPhoto))
	return NewImageWriter(w.Context, w.Request, w.ReplyID != 0)
}

func NewImageWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *TelegramWriter {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Telegram.API,
		ChatID:  request.GetMessage().Source.Telegram.Update.Message.Chat.ID,
		Type:    "image",
		Request: request,
		bufs:    []bytes.Buffer{},
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

func (w *TelegramWriter) closeImage() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	bufs := w.bufs
	if w.mainWriter != nil {
		bufs = w.mainWriter.bufs
		w.mainWriter.closedBufs++
		if w.mainWriter.closedBufs != len(w.mainWriter.bufs) {
			return nil
		}
	}

	logger.Debug().Int64("chatID", w.ChatID).Msg("replying with image")

	if w.mainWriter == nil {
		msg := tgbotapi.NewPhoto(w.ChatID, tgbotapi.FileReader{
			Name:   "generated.png",
			Reader: bytes.NewReader(bufs[0].Bytes()),
		})
		msg.ReplyToMessageID = w.ReplyID
		_, err := w.Bot.Send(msg)
		return err
	}

	var files []interface{}
	for k := range bufs {
		files = append(files, tgbotapi.NewInputMediaPhoto(
			tgbotapi.FileReader{
				Name:   "generated.png",
				Reader: bytes.NewReader(bufs[k].Bytes()),
			},
		))
	}

	msg := tgbotapi.NewMediaGroup(w.ChatID, files)
	msg.ReplyToMessageID = w.ReplyID
	_, err := w.Bot.SendMediaGroup(msg)
	return err
}
