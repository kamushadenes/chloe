package telegram

import (
	"context"
	"net/http"
	"time"

	"github.com/aquilax/truncate"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

type TelegramWriter struct {
	Context          context.Context
	Prompt           string
	Bot              *tgbotapi.BotAPI
	ChatID           int64
	Type             string
	ReplyID          int
	Request          structs.ActionOrCompletionRequest
	objs             []*response_object_structs.ResponseObject
	externalID       int
	lastUpdate       *time.Time
	preWriteCallback func()
}

func NewTelegramWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *TelegramWriter {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Telegram.API,
		ChatID:  request.GetMessage().Source.Telegram.Update.Message.Chat.ID,
		Request: request,
	}

	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Telegram.Update.Message.MessageID
	}

	return w
}

func (w *TelegramWriter) Flush() {
	if w.Type != "text" {
		return
	}

	if w.externalID == 0 && (config.Telegram.SendProcessingMessage || config.Telegram.StreamMessages) {
		msg, err := w.Bot.Send(tgbotapi.NewMessage(w.ChatID, "↻ Processing..."))
		if err != nil {
			return
		}
		w.externalID = msg.MessageID
		tt := time.Now()
		w.lastUpdate = &tt
	}

	if !config.Telegram.StreamMessages {
		return
	}

	if w.lastUpdate != nil && time.Since(*w.lastUpdate) > config.Telegram.StreamFlushInterval {
		_, _ = w.Bot.Send(tgbotapi.NewEditMessageText(w.ChatID, w.externalID, truncate.Truncate(
			string(w.objs[0].Data),
			config.Telegram.MaxMessageLength,
			"...",
			truncate.PositionEnd,
		)))
		tt := time.Now()
		w.lastUpdate = &tt
	}
}

func (w *TelegramWriter) Close() error {
	funcs := []func() error{
		w.closeText,
		w.closeAudio,
		w.closeImage,
	}

	for k := range funcs {
		if err := funcs[k](); err != nil {
			return err
		}
	}

	return nil
}

func (w *TelegramWriter) Write(p []byte) (n int, err error) {
	if w.preWriteCallback != nil {
		w.preWriteCallback()
	}

	if len(w.objs) == 0 {
		w.objs = append(w.objs, &response_object_structs.ResponseObject{
			Type:   response_object_structs.Text,
			Result: true,
		})
	}

	w.objs[0].Data = append(w.objs[0].Data, p...)

	return len(p), nil
}

func (w *TelegramWriter) WriteObject(obj *response_object_structs.ResponseObject) error {
	w.objs = append(w.objs, obj)

	return nil
}

func (w *TelegramWriter) SetPrompt(prompt string) {
	w.Prompt = prompt
}

func (w *TelegramWriter) WriteHeader(statusCode int) {}
func (w *TelegramWriter) Header() http.Header        { return http.Header{} }
func (w *TelegramWriter) SetPreWriteCallback(fn func()) {
	w.preWriteCallback = fn
}

func (w *TelegramWriter) GetObjects() []*response_object_structs.ResponseObject {
	return w.objs
}
