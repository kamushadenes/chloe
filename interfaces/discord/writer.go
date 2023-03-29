package discord

import (
	"bytes"
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"time"
)

type DiscordWriter struct {
	Context    context.Context
	Prompt     string
	Bot        *discordgo.Session
	ChatID     string
	Type       string
	ReplyID    string
	Request    structs.ActionOrCompletionRequest
	bufs       []bytes.Buffer
	bufID      int
	closedBufs int
	mainWriter *DiscordWriter
	externalID string
	lastUpdate *time.Time
}

func (w *DiscordWriter) Flush() {
	if w.Type != "text" {
		return
	}

	if w.externalID == "" && (config.Discord.SendProcessingMessage || config.Discord.StreamMessages) {
		msg, err := w.Bot.ChannelMessageSend(w.ChatID, "↻ Processing...")
		if err != nil {
			return
		}
		w.externalID = msg.ID
		tt := time.Now()
		w.lastUpdate = &tt
	}

	if !config.Discord.StreamMessages {
		return
	}

	if w.lastUpdate != nil && time.Now().Sub(*w.lastUpdate) > config.Discord.StreamFlushInterval {
		_, _ = w.Bot.ChannelMessageEdit(w.ChatID, w.externalID, w.bufs[0].String()[:config.Discord.MaxMessageLength])
		tt := time.Now()
		w.lastUpdate = &tt
	}
}

func (w *DiscordWriter) Close() error {
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

func (w *DiscordWriter) Write(p []byte) (n int, err error) {
	if w.mainWriter != nil {
		return w.mainWriter.bufs[w.bufID].Write(p)
	}

	return w.bufs[0].Write(p)
}

func (w *DiscordWriter) Subwriter() *DiscordWriter {
	w.bufs = append(w.bufs, bytes.Buffer{})

	return &DiscordWriter{
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

func (w *DiscordWriter) SetPrompt(prompt string) {
	w.Prompt = prompt
}
