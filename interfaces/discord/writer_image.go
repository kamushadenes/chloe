package discord

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
)

func (w *DiscordWriter) ToImageWriter() *DiscordWriter {
	return NewImageWriter(w.Context, w.Request, w.ReplyID != "", w.Prompt)
}

func NewImageWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *DiscordWriter {
	w := &DiscordWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Discord.API,
		ChatID:  request.GetMessage().Source.Discord.Message.ChannelID,
		Type:    "image",
		Request: request,
		bufs:    []bytes.Buffer{},
		bufID:   0,
	}
	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Discord.Message.ID
	}

	return w
}

func (w *DiscordWriter) closeImage() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	bufs := w.bufs

	if w.mainWriter != nil {
		bufs = w.mainWriter.bufs
		w.mainWriter.closedBufs++
		if w.mainWriter.closedBufs != len(w.mainWriter.bufs) {
			return nil
		}
	}

	logger.Debug().Str("chatID", w.ChatID).Msg("replying with image")

	var files []*discordgo.File
	for k := range bufs {
		files = append(files, &discordgo.File{
			Name:        fmt.Sprintf("generated-%d.png", k),
			ContentType: "image/png",
			Reader:      bytes.NewReader(bufs[k].Bytes()),
		})
	}

	content := fmt.Sprintf("Prompt: %s", w.Prompt)

	_, err := w.Bot.ChannelMessageSendComplex(w.ChatID, &discordgo.MessageSend{
		Files:   files,
		Content: content,
	})
	return err
}
