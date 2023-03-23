package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
)

func tryAndRespond(ctx context.Context, msg *memory.Message, successText, errorText string, err error, reply bool) {
	logger := zerolog.Ctx(ctx)

	text := successText

	if err != nil {
		logger.Error().Err(err).Msg("error processing")
		text = errorText
	}

	if text != "" {
		tmsg := tgbotapi.NewMessage(msg.Source.Telegram.Update.Message.Chat.ID, text)
		tmsg.ParseMode = tgbotapi.ModeMarkdownV2
		if reply {
			tmsg.ReplyToMessageID = msg.Source.Telegram.Update.Message.MessageID
		}
		if _, err := msg.Source.Telegram.API.Send(tmsg); err != nil {
			logger.Error().Err(err).Msg("error sending message")
		}
	}
}

func userFromMessage(ctx context.Context, msg *memory.Message) (*memory.User, error) {
	return memory.GetUserByExternalID(ctx, fmt.Sprintf("%d", msg.Source.Telegram.Update.Message.From.ID), "telegram")
}

func promptFromMessage(msg *memory.Message) string {
	if msg.Source.Telegram.Update.Message.CommandArguments() != "" {
		return msg.Source.Telegram.Update.Message.CommandArguments()
	} else {
		return msg.Source.Telegram.Update.Message.Text
	}
}
