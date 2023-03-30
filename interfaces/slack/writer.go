package slack

import (
	"bytes"
	"context"
	"github.com/aquilax/truncate"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"github.com/slack-go/slack"
	"time"
)

type SlackWriter struct {
	Context    context.Context
	Prompt     string
	Bot        *slack.Client
	ChatID     string
	Type       string
	ReplyID    string
	Request    structs.ActionOrCompletionRequest
	bufs       []bytes.Buffer
	bufID      int
	closedBufs int
	mainWriter *SlackWriter
	externalID string
	lastUpdate *time.Time
}

func (w *SlackWriter) Flush() {
	if w.Type != "text" {
		return
	}

	if w.externalID == "" && (config.Slack.SendProcessingMessage || config.Slack.StreamMessages) {
		_, ts, err := w.Bot.PostMessage(w.ChatID, slack.MsgOptionText("↻ Processing...", false))
		if err != nil {
			return
		}
		w.externalID = ts
		tt := time.Now()
		w.lastUpdate = &tt
	}

	if !config.Slack.StreamMessages {
		return
	}

	if w.lastUpdate != nil && time.Since(*w.lastUpdate) > config.Slack.StreamFlushInterval {
		_, _, _, _ = w.Bot.UpdateMessage(w.ChatID, w.externalID, slack.MsgOptionText(
			truncate.Truncate(
				w.bufs[0].String(),
				config.Slack.MaxMessageLength,
				"...",
				truncate.PositionEnd,
			), false))

		tt := time.Now()
		w.lastUpdate = &tt
	}
}

func (w *SlackWriter) Close() error {
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

func (w *SlackWriter) Write(p []byte) (n int, err error) {
	if w.mainWriter != nil {
		return w.mainWriter.bufs[w.bufID].Write(p)
	}

	return w.bufs[0].Write(p)
}

func (w *SlackWriter) Subwriter() *SlackWriter {
	w.bufs = append(w.bufs, bytes.Buffer{})

	return &SlackWriter{
		Context:    w.Context,
		Bot:        w.Bot,
		ChatID:     w.ChatID,
		Type:       w.Type,
		ReplyID:    w.ReplyID,
		Request:    w.Request,
		Prompt:     w.Prompt,
		bufs:       []bytes.Buffer{{}},
		bufID:      len(w.bufs) - 1,
		mainWriter: w,
	}
}

func (w *SlackWriter) SetPrompt(prompt string) {
	w.Prompt = prompt
}
