package telegram

import (
	"bytes"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type TelegramWriter struct {
	Context    context.Context
	Bot        *tgbotapi.BotAPI
	ChatID     int64
	Type       string
	ReplyID    int
	msg        *memory.Message
	bufs       []bytes.Buffer
	bufId      int
	closedBufs int
	mainWriter *TelegramWriter
}

func (t *TelegramWriter) Close() error {
	switch t.Type {
	case "text":
		if t.bufs[0].Len() == 0 {
			return nil
		}
		msg := tgbotapi.NewMessage(t.ChatID, t.bufs[0].String())

		msg.ParseMode = tgbotapi.ModeMarkdown
		msg.ReplyToMessageID = t.ReplyID
		_, err := t.Bot.Send(msg)
		if err != nil {
			msg.ParseMode = ""
			_, err = t.Bot.Send(msg)
		}
		return err
	case "audio":
		tmsg := tgbotapi.NewVoice(t.ChatID, tgbotapi.FileReader{
			Reader: bytes.NewReader(t.bufs[0].Bytes()),
		})
		tmsg.ReplyToMessageID = t.ReplyID
		_, err := t.Bot.Send(tmsg)
		return err
	case "image":
		if t.mainWriter == nil {
			msg := tgbotapi.NewPhoto(t.ChatID, tgbotapi.FileReader{
				Reader: bytes.NewReader(t.bufs[0].Bytes()),
			})
			msg.ReplyToMessageID = t.ReplyID
			_, err := t.Bot.Send(msg)
			return err
		} else {
			t.mainWriter.closedBufs++
			if t.mainWriter.closedBufs != len(t.mainWriter.bufs) {
				return nil
			}
		}

		var files []interface{}
		for _, buf := range t.mainWriter.bufs {
			files = append(files, tgbotapi.NewInputMediaPhoto(
				tgbotapi.FileReader{
					Reader: bytes.NewReader(buf.Bytes()),
				},
			))
		}
		msg := tgbotapi.NewMediaGroup(t.ChatID, files)
		msg.ReplyToMessageID = t.ReplyID
		_, err := t.Bot.SendMediaGroup(msg)
		return err
	}
	return nil
}

func (t *TelegramWriter) Write(p []byte) (n int, err error) {
	if t.mainWriter != nil {
		return t.mainWriter.bufs[t.bufId].Write(p)
	}

	return t.bufs[0].Write(p)
}

func (t *TelegramWriter) Subwriter() *TelegramWriter {
	t.bufs = append(t.bufs, bytes.Buffer{})

	return &TelegramWriter{
		Context:    t.Context,
		Bot:        t.Bot,
		ChatID:     t.ChatID,
		Type:       t.Type,
		ReplyID:    t.ReplyID,
		bufs:       []bytes.Buffer{{}},
		bufId:      len(t.bufs) - 1,
		mainWriter: t,
	}
}

func (t *TelegramWriter) ToImageWriter() io.WriteCloser {
	t.Bot.Send(tgbotapi.NewChatAction(t.ChatID, tgbotapi.ChatUploadPhoto))
	return NewImageWriter(t.Context, t.msg, t.ReplyID != 0)
}

func (t *TelegramWriter) ToTextWriter() io.WriteCloser {
	return NewTextWriter(t.Context, t.msg, t.ReplyID != 0)
}

func (t *TelegramWriter) ToAudioWriter() io.WriteCloser {
	t.Bot.Send(tgbotapi.NewChatAction(t.ChatID, tgbotapi.ChatRecordVoice))
	return NewAudioWriter(t.Context, t.msg, t.ReplyID != 0)
}

func NewTextWriter(ctx context.Context, msg *memory.Message, reply bool) io.WriteCloser {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     msg.Source.Telegram.API,
		ChatID:  msg.Source.Telegram.Update.Message.Chat.ID,
		Type:    "text",
		bufs:    []bytes.Buffer{{}},
		bufId:   0,
		msg:     msg,
	}

	if reply {
		w.ReplyID = msg.Source.Telegram.Update.Message.MessageID
	}

	return w
}

func NewImageWriter(ctx context.Context, msg *memory.Message, reply bool) io.WriteCloser {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     msg.Source.Telegram.API,
		ChatID:  msg.Source.Telegram.Update.Message.Chat.ID,
		Type:    "image",
		bufs:    []bytes.Buffer{},
		bufId:   0,
		msg:     msg,
	}

	if reply {
		w.ReplyID = msg.Source.Telegram.Update.Message.MessageID
	}

	return w
}

func NewAudioWriter(ctx context.Context, msg *memory.Message, reply bool) io.WriteCloser {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     msg.Source.Telegram.API,
		ChatID:  msg.Source.Telegram.Update.Message.Chat.ID,
		Type:    "audio",
		bufs:    []bytes.Buffer{{}},
		bufId:   0,
		msg:     msg,
	}

	if reply {
		w.ReplyID = msg.Source.Telegram.Update.Message.MessageID
	}

	return w
}
