package telegram

import (
	"bytes"
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"io"
)

type TelegramWriter struct {
	Context    context.Context
	Bot        *tgbotapi.BotAPI
	ChatID     int64
	Type       string
	ReplyID    int
	Request    structs.Request
	bufs       []bytes.Buffer
	bufId      int
	closedBufs int
	mainWriter *TelegramWriter
}

func (t *TelegramWriter) Close() error {
	logger := zerolog.Ctx(t.Context).With().Str("requestID", t.Request.GetID()).Logger()

	switch t.Type {
	case "text":
		if t.bufs[0].Len() == 0 {
			return nil
		}
		logger.Debug().Int64("chatID", t.ChatID).Msg("replying with text")

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
		logger.Debug().Int64("chatID", t.ChatID).Msg("replying with audio")
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

		logger.Debug().Int64("chatID", t.ChatID).Msg("replying with image")

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
		Request:    t.Request,
		bufs:       []bytes.Buffer{{}},
		bufId:      len(t.bufs) - 1,
		mainWriter: t,
	}
}

func (t *TelegramWriter) ToImageWriter() io.WriteCloser {
	_, _ = t.Bot.Send(tgbotapi.NewChatAction(t.ChatID, tgbotapi.ChatUploadPhoto))
	return NewImageWriter(t.Context, t.Request, t.ReplyID != 0)
}

func (t *TelegramWriter) ToTextWriter() io.WriteCloser {
	_, _ = t.Bot.Send(tgbotapi.NewChatAction(t.ChatID, tgbotapi.ChatTyping))
	return NewTextWriter(t.Context, t.Request, t.ReplyID != 0)
}

func (t *TelegramWriter) ToAudioWriter() io.WriteCloser {
	_, _ = t.Bot.Send(tgbotapi.NewChatAction(t.ChatID, tgbotapi.ChatRecordVoice))
	return NewAudioWriter(t.Context, t.Request, t.ReplyID != 0)
}

func NewTextWriter(ctx context.Context, request structs.Request, reply bool) io.WriteCloser {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Telegram.API,
		ChatID:  request.GetMessage().Source.Telegram.Update.Message.Chat.ID,
		Type:    "text",
		Request: request,
		bufs:    []bytes.Buffer{{}},
		bufId:   0,
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Telegram.Update.Message.MessageID
	}

	return w
}

func NewImageWriter(ctx context.Context, request structs.Request, reply bool) io.WriteCloser {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Telegram.API,
		ChatID:  request.GetMessage().Source.Telegram.Update.Message.Chat.ID,
		Type:    "image",
		Request: request,
		bufs:    []bytes.Buffer{},
		bufId:   0,
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Telegram.Update.Message.MessageID
	}

	return w
}

func NewAudioWriter(ctx context.Context, request structs.Request, reply bool) io.WriteCloser {
	w := &TelegramWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Telegram.API,
		ChatID:  request.GetMessage().Source.Telegram.Update.Message.Chat.ID,
		Type:    "audio",
		Request: request,
		bufs:    []bytes.Buffer{{}},
		bufId:   0,
	}

	if reply {
		w.ReplyID = request.GetMessage().Source.Telegram.Update.Message.MessageID
	}

	return w
}
