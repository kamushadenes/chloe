package discord

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type DiscordWriter struct {
	Context    context.Context
	Prompt     string
	Bot        *discordgo.Session
	ChatID     string
	Type       string
	ReplyID    string
	Request    structs.Request
	bufs       []bytes.Buffer
	bufId      int
	closedBufs int
	mainWriter *DiscordWriter
}

func (t *DiscordWriter) Close() error {
	switch t.Type {
	case "text":
		return t.Request.GetMessage().SendText(t.bufs[0].String())
	case "audio":
		msg := &discordgo.MessageSend{
			File: &discordgo.File{
				Name:        "generated.mp3",
				ContentType: "audio/mpeg",
				Reader:      bytes.NewReader(t.bufs[0].Bytes()),
			},
			Content: fmt.Sprintf("Prompt: %s", t.Prompt),
		}
		_, err := t.Bot.ChannelMessageSendComplex(t.ChatID, msg)
		return err
	case "image":
		if t.mainWriter == nil {
			msg := &discordgo.MessageSend{
				File: &discordgo.File{
					Name:        "generated.png",
					ContentType: "image/png",
					Reader:      bytes.NewReader(t.bufs[0].Bytes()),
				},
				Content: fmt.Sprintf("Prompt: %s", t.Prompt),
			}
			_, err := t.Bot.ChannelMessageSendComplex(t.ChatID, msg)
			return err
		} else {
			t.mainWriter.closedBufs++
			if t.mainWriter.closedBufs != len(t.mainWriter.bufs) {
				return nil
			}
		}

		var files []*discordgo.File
		for _, buf := range t.mainWriter.bufs {
			files = append(files, &discordgo.File{
				Name:        "generated.png",
				ContentType: "image/png",
				Reader:      bytes.NewReader(buf.Bytes()),
			})
		}
		msg := &discordgo.MessageSend{
			Files:   files,
			Content: fmt.Sprintf("Prompt: %s", t.Prompt),
		}
		_, err := t.Bot.ChannelMessageSendComplex(t.ChatID, msg)
		return err
	}

	return nil
}

func (t *DiscordWriter) Write(p []byte) (n int, err error) {
	if t.mainWriter != nil {
		return t.mainWriter.bufs[t.bufId].Write(p)
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
		bufId:      len(t.bufs) - 1,
		mainWriter: t,
	}
}

func (t *DiscordWriter) SetPrompt(prompt string) {
	t.Prompt = prompt
}

func (t *DiscordWriter) ToImageWriter() io.WriteCloser {
	return NewImageWriter(t.Context, t.Request, t.ReplyID != "", t.Prompt)
}

func (t *DiscordWriter) ToTextWriter() io.WriteCloser {
	return NewTextWriter(t.Context, t.Request, t.ReplyID != "", t.Prompt)
}

func (t *DiscordWriter) ToAudioWriter() io.WriteCloser {
	return NewAudioWriter(t.Context, t.Request, t.ReplyID != "", t.Prompt)
}

func NewTextWriter(ctx context.Context, request structs.Request, reply bool, prompt ...string) io.WriteCloser {
	w := &DiscordWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Discord.API,
		ChatID:  request.GetMessage().Source.Discord.Message.ChannelID,
		Type:    "text",
		Request: request,
		bufs:    []bytes.Buffer{{}},
		bufId:   0,
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Discord.Message.ID
	}
	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	return w
}

func NewImageWriter(ctx context.Context, request structs.Request, reply bool, prompt ...string) io.WriteCloser {
	w := &DiscordWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Discord.API,
		ChatID:  request.GetMessage().Source.Discord.Message.ChannelID,
		Type:    "image",
		Request: request,
		bufs:    []bytes.Buffer{},
		bufId:   0,
	}
	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Discord.Message.ID
	}

	return w
}

func NewAudioWriter(ctx context.Context, request structs.Request, reply bool, prompt ...string) io.WriteCloser {
	w := &DiscordWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Discord.API,
		ChatID:  request.GetMessage().Source.Discord.Message.ChannelID,
		Type:    "audio",
		Request: request,
		bufs:    []bytes.Buffer{{}},
		bufId:   0,
	}
	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Discord.Message.ID
	}

	return w
}
