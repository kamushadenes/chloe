package discord

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
	"net/http"
	"time"

	"github.com/aquilax/truncate"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

type DiscordWriter struct {
	Context          context.Context
	Prompt           string
	Bot              *discordgo.Session
	ChatID           string
	Type             string
	ReplyID          string
	Message          *memory.Message
	objs             []*response_object_structs.ResponseObject
	externalID       string
	lastUpdate       *time.Time
	preWriteCallback func()
}

func NewDiscordWriter(ctx context.Context, msg *memory.Message, reply bool, prompt ...string) *DiscordWriter {
	w := &DiscordWriter{
		Context: ctx,
		Bot:     msg.Source.Discord.API,
		ChatID:  msg.Source.Discord.Message.ChannelID,
		Message: msg,
	}

	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	if reply {
		w.ReplyID = msg.Source.Discord.Message.ID
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
			logger.Error().Err(err).Msgf("error closing writer %d", k)
			return err
		}
	}

	return nil
}

func (w *DiscordWriter) Write(p []byte) (n int, err error) {
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

func (w *DiscordWriter) WriteObject(obj *response_object_structs.ResponseObject) error {
	w.objs = append(w.objs, obj)

	return nil
}

func (w *DiscordWriter) SetPrompt(prompt string) {
	w.Prompt = prompt
}

func (w *DiscordWriter) WriteHeader(statusCode int) {}
func (w *DiscordWriter) Header() http.Header        { return http.Header{} }
func (w *DiscordWriter) SetPreWriteCallback(fn func()) {
	w.preWriteCallback = fn
}

func (w *DiscordWriter) GetObjects() []*response_object_structs.ResponseObject {
	return w.objs
}
