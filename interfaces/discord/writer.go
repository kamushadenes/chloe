package discord

import (
	"context"
	"github.com/aquilax/truncate"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
	"net/http"
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
	objs       []*structs.ResponseObject
	externalID string
	lastUpdate *time.Time
}

func NewDiscordWriter(ctx context.Context, req structs.ActionOrCompletionRequest, reply bool, prompt ...string) *DiscordWriter {
	w := &DiscordWriter{
		Context: ctx,
		Bot:     req.GetMessage().Source.Discord.API,
		ChatID:  req.GetMessage().Source.Discord.Message.ChannelID,
		Request: req,
	}

	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	if reply {
		w.ReplyID = req.GetMessage().Source.Discord.Message.ID
	}

	return w
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

	if w.lastUpdate != nil && time.Since(*w.lastUpdate) > config.Discord.StreamFlushInterval {
		_, _ = w.Bot.ChannelMessageEdit(w.ChatID, w.externalID, truncate.Truncate(
			w.objs[0].String(),
			config.Discord.MaxMessageLength,
			"...",
			truncate.PositionEnd,
		))
		tt := time.Now()
		w.lastUpdate = &tt
	}
}

func (w *DiscordWriter) Close() error {
	logger := logging.FromContext(w.Context)

	funcs := []func() error{
		w.closeText,
		w.closeAudio,
		w.closeImage,
	}

	for k := range funcs {
		if err := funcs[k](); err != nil {
			logger.Error().Err(err).Str("requestID", w.Request.GetID()).Msgf("error closing writer %d", k)
			return err
		}
	}

	return nil
}

func (w *DiscordWriter) Write(p []byte) (n int, err error) {
	if len(w.objs) == 0 {
		w.objs = append(w.objs, &structs.ResponseObject{
			Type:   structs.Text,
			Result: true,
		})
	}

	w.objs[0].Data = append(w.objs[0].Data, p...)

	return len(p), nil
}

func (w *DiscordWriter) WriteObject(obj *structs.ResponseObject) error {
	w.objs = append(w.objs, obj)

	return nil
}

func (w *DiscordWriter) SetPrompt(prompt string) {
	w.Prompt = prompt
}

func (w *DiscordWriter) WriteHeader(statusCode int) {}
func (w *DiscordWriter) Header() http.Header        { return nil }
