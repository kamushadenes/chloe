package discord

import (
	"bytes"
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
)

func (w *DiscordWriter) ToTextWriter() *DiscordWriter {
	return NewTextWriter(w.Context, w.Request, w.ReplyID != "", w.Prompt)
}

func NewTextWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *DiscordWriter {
	w := &DiscordWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Discord.API,
		ChatID:  request.GetMessage().Source.Discord.Message.ChannelID,
		Type:    "text",
		Request: request,
		bufs:    []bytes.Buffer{{}},
		bufID:   0,
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Discord.Message.ID
	}
	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	return w
}

func (w *DiscordWriter) closeText() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	logger.Debug().Str("chatID", w.ChatID).Msg("replying with text")

	msgs := utils.StringToChunks(w.bufs[0].String(), config.Discord.MaxMessageLength)

	if w.externalID == "" {
		for k := range msgs {
			if err := w.Request.GetMessage().SendText(msgs[k], true); err != nil {
				return err
			}
		}
	} else {
		if _, err := w.Bot.ChannelMessageEdit(w.ChatID, w.externalID, msgs[0]); err != nil {
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
