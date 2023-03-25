package telegram

import (
	"context"
	"fmt"
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

	if reply {
		err = msg.SendText(text, msg.Source.Telegram.Update.Message.MessageID)
	} else {
		err = msg.SendText(text)
	}

	if err != nil {
		logger.Error().Err(err).Msg("error sending message")
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
