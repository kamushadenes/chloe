package discord

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
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

func (t *DiscordWriter) Flush() {
	if t.Type != "text" {
		return
	}

	if t.externalID == "" && (config.Discord.SendProcessingMessage || config.Discord.StreamMessages) {
		msg, err := t.Bot.ChannelMessageSend(t.ChatID, "↻ Processing...")
		if err != nil {
			return
		}
		t.externalID = msg.ID
		tt := time.Now()
		t.lastUpdate = &tt
	}

	if !config.Discord.StreamMessages {
		return
	}

	if t.lastUpdate != nil && time.Now().Sub(*t.lastUpdate) > config.Discord.StreamFlushInterval {
		_, _ = t.Bot.ChannelMessageEdit(t.ChatID, t.externalID, t.bufs[0].String()[:config.Discord.MaxMessageLength])
		tt := time.Now()
		t.lastUpdate = &tt
	}
}

func (t *DiscordWriter) Close() error {
	logger := zerolog.Ctx(t.Context).With().Str("requestID", t.Request.GetID()).Logger()

	switch t.Type {
	case "text":
		logger.Debug().Str("chatID", t.ChatID).Msg("replying with text")

		msgs := utils.StringToChunks(t.bufs[0].String(), config.Discord.MaxMessageLength)

		if t.externalID == "" {
			for k := range msgs {
				if err := t.Request.GetMessage().SendText(msgs[k], true); err != nil {
					return err
				}
			}
		} else {
			if _, err := t.Bot.ChannelMessageEdit(t.ChatID, t.externalID, msgs[0]); err != nil {
				return err
			}

			for k := range msgs[1:] {
				if err := t.Request.GetMessage().SendText(msgs[k], true); err != nil {
					return err
				}
			}
		}
	case "audio":
		logger.Debug().Str("chatID", t.ChatID).Msg("replying with audio")
		bufs := t.bufs

		if t.mainWriter != nil {
			bufs = t.mainWriter.bufs
			t.mainWriter.closedBufs++
			if t.mainWriter.closedBufs != len(t.mainWriter.bufs) {
				return nil
			}
		}

		var files []*discordgo.File
		for k := range bufs {
			files = append(files, &discordgo.File{
				Name:        "generated.mp3",
				ContentType: "audio/mpeg",
				Reader:      bytes.NewReader(bufs[k].Bytes()),
			})
		}

		_, err := t.Bot.ChannelMessageSendComplex(t.ChatID, &discordgo.MessageSend{
			Files:   files,
			Content: fmt.Sprintf("Prompt: %s", t.Prompt),
		})
		return err
	case "image":
		bufs := t.bufs

		if t.mainWriter != nil {
			bufs = t.mainWriter.bufs
			t.mainWriter.closedBufs++
			if t.mainWriter.closedBufs != len(t.mainWriter.bufs) {
				return nil
			}
		}

		logger.Debug().Str("chatID", t.ChatID).Msg("replying with image")

		var files []*discordgo.File
		for k := range bufs {
			files = append(files, &discordgo.File{
				Name:        "generated.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(bufs[k].Bytes()),
			})
		}

		content := fmt.Sprintf("Prompt: %s", t.Prompt)

		_, err := t.Bot.ChannelMessageSendComplex(t.ChatID, &discordgo.MessageSend{
			Files:   files,
			Content: content,
		})
		return err
	}

	return nil
}

func (t *DiscordWriter) Write(p []byte) (n int, err error) {
	if t.mainWriter != nil {
		return t.mainWriter.bufs[t.bufID].Write(p)
	}

	return t.bufs[0].Write(p)
}

func (t *DiscordWriter) Subwriter() *DiscordWriter {
	t.bufs = append(t.bufs, bytes.Buffer{})

	return &DiscordWriter{
		Context:    t.Context,
		Bot:        t.Bot,
		ChatID:     t.ChatID,
		Type:       t.Type,
		ReplyID:    t.ReplyID,
		Request:    t.Request,
		Prompt:     t.Prompt,
		bufs:       []bytes.Buffer{{}},
		bufID:      len(t.bufs) - 1,
		mainWriter: t,
	}
}

func (t *DiscordWriter) SetPrompt(prompt string) {
	t.Prompt = prompt
}

func (t *DiscordWriter) ToImageWriter() *DiscordWriter {
	return NewImageWriter(t.Context, t.Request, t.ReplyID != "", t.Prompt)
}

func (t *DiscordWriter) ToTextWriter() *DiscordWriter {
	return NewTextWriter(t.Context, t.Request, t.ReplyID != "", t.Prompt)
}

func (t *DiscordWriter) ToAudioWriter() *DiscordWriter {
	return NewAudioWriter(t.Context, t.Request, t.ReplyID != "", t.Prompt)
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

func NewAudioWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *DiscordWriter {
	w := &DiscordWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Discord.API,
		ChatID:  request.GetMessage().Source.Discord.Message.ChannelID,
		Type:    "audio",
		Request: request,
		bufs:    []bytes.Buffer{{}},
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
