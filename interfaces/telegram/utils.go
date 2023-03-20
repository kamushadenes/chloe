package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/messages"
	"github.com/kamushadenes/chloe/users"
	"github.com/rs/zerolog"
)

func tryAndRespond(ctx context.Context, msg *messages.Message, successText, errorText string, err error, reply bool) {
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

func userFromMessage(ctx context.Context, msg *messages.Message) (*users.User, error) {
	return users.GetUserOrSetByExternalId(ctx, fmt.Sprintf("%d", msg.Source.Telegram.Update.Message.From.ID),
		"telegram",
		msg.Source.Telegram.Update.Message.From.FirstName,
		msg.Source.Telegram.Update.Message.From.LastName,
		msg.Source.Telegram.Update.Message.From.UserName)
}

func promptFromMessage(msg *messages.Message) string {
	if msg.Source.Telegram.Update.Message.CommandArguments() != "" {
		return msg.Source.Telegram.Update.Message.CommandArguments()
	} else {
		return msg.Source.Telegram.Update.Message.Text
	}
}
