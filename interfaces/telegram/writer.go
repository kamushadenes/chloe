package telegram

import (
	"bytes"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"time"
)

type TelegramWriter struct {
	Context    context.Context
	Prompt     string
	Bot        *tgbotapi.BotAPI
	ChatID     int64
	Type       string
	ReplyID    int
	Request    structs.ActionOrCompletionRequest
	bufs       []bytes.Buffer
	bufID      int
	closedBufs int
	mainWriter *TelegramWriter
	externalID int
	lastUpdate *time.Time
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

	if w.lastUpdate != nil && time.Now().Sub(*w.lastUpdate) > config.Telegram.StreamFlushInterval {
		_, _ = w.Bot.Send(tgbotapi.NewEditMessageText(w.ChatID, w.externalID, w.bufs[0].String()[:config.Telegram.MaxMessageLength]))
		tt := time.Now()
		w.lastUpdate = &tt
	}
}

func (w *TelegramWriter) Close() error {
	if (len(w.bufs) > 0 && w.bufs[0].Len() > 0) ||
		(w.mainWriter != nil && len(w.mainWriter.bufs) > 0 && w.mainWriter.bufs[0].Len() > 0) {
		switch w.Type {
		case "text":
			return w.closeText()
		case "audio":
			return w.closeAudio()
		case "image":
			return w.closeImage()
		}
	}

	return nil
}

func (w *TelegramWriter) Write(p []byte) (n int, err error) {
	if w.mainWriter != nil {
		return w.mainWriter.bufs[w.bufID].Write(p)
	}

	return w.bufs[0].Write(p)
}

func (w *TelegramWriter) SetPrompt(prompt string) {
	w.Prompt = prompt
}

func (w *TelegramWriter) Subwriter() *TelegramWriter {
	w.bufs = append(w.bufs, bytes.Buffer{})

	return &TelegramWriter{
		Context:    w.Context,
		Prompt:     w.Prompt,
		Bot:        w.Bot,
		ChatID:     w.ChatID,
		Type:       w.Type,
		ReplyID:    w.ReplyID,
		Request:    w.Request,
		bufs:       []bytes.Buffer{{}},
		bufID:      len(w.bufs) - 1,
		mainWriter: w,
	}
}
