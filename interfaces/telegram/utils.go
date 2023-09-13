package telegram

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
)

func tryAndRespond(ctx context.Context, msg *memory.Message, successText, errorText string, err error, reply bool) {
	logger := logging.FromContext(ctx)

	text := successText

	if err != nil {
		logger.Error().Err(err).Msg("error processing")
		text = errorText
	}

	if reply {
		err = msg.SendText(text, true, msg.Source.Telegram.Update.Message.MessageID)
	} else {
		err = msg.SendText(text, true)
	}

	if err != nil {
		logger.Error().Err(err).Msg("error sending message")
	}
}

func userFromMessage(ctx context.Context, msg *memory.Message) (*memory.User, error) {
	logger := logging.FromContext(ctx)

	user, err := memory.GetUserByExternalID(ctx, fmt.Sprintf("%d", msg.Source.Telegram.Update.Message.From.ID), "telegram")
	if err != nil {
		user, err = memory.CreateUser(ctx, msg.Source.Telegram.Update.Message.From.FirstName, msg.Source.Telegram.Update.Message.From.LastName, msg.Source.Telegram.Update.Message.From.UserName)
		if err != nil {
			logger.Error().Err(err).Msg("error getting user from message")
			return nil, err
		}
		err = user.AddExternalID(ctx, fmt.Sprintf("%d", msg.Source.Telegram.Update.Message.From.ID), "telegram")
		if err != nil {
			logger.Error().Err(err).Msg("error getting user from message")
			return nil, err
		}
	}

	return user, err
}

func promptFromMessage(msg *memory.Message) string {
	if msg.Source.Telegram.Update.Message.CommandArguments() != "" {
		return msg.Source.Telegram.Update.Message.CommandArguments()
	} else {
		return msg.Source.Telegram.Update.Message.Text
	}
}
