package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
)

func handleTextUpdate(ctx context.Context, msg *memory.Message) {
	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
	_ = processText(ctx, msg, nil)
}

func handleAudioUpdate(ctx context.Context, msg *memory.Message) {
	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))

	if err := processAudio(ctx, msg); err != nil {
		_ = msg.SendError(err)
	}
}

func handleImageUpdate(ctx context.Context, msg *memory.Message) {
	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatUploadPhoto))

	if err := processImage(ctx, msg); err != nil {
		_ = msg.SendError(err)
	}
}

func handleUpdates(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger := zerolog.Ctx(ctx)
	logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str("externalUserID", fmt.Sprintf("%d", update.Message.From.ID))
	})

	if update.Message == nil { // ignore non-Message updates
		return
	}

	msg := memory.NewMessage(fmt.Sprintf("%d", update.Message.MessageID), "telegram")
	msg.Role = "user"
	msg.SetContent(update.Message.Text)
	msg.Source.Telegram = &memory.TelegramMessageSource{
		API:    bot,
		Update: update,
	}

	user, err := userFromMessage(ctx, msg)
	if err != nil {
		logger.Error().Err(err).Msg("error getting user from message")
		_ = msg.SendError(err)
		return
	}
	msg.User = user

	if err := channels.RegisterIncomingMessage(msg); err != nil {
		_ = msg.SendError(err)
		return
	}

	if handleCommands(ctx, msg) {
		return
	}

	if update.Message.Text != "" {
		handleTextUpdate(ctx, msg)
	}

	if update.Message.Voice != nil || update.Message.Audio != nil {
		if update.Message.Voice != nil {
			msg.AddAudio(downloadFile(ctx, bot, update.Message.Voice.FileID))
		}

		if update.Message.Audio != nil {
			msg.AddAudio(downloadFile(ctx, bot, update.Message.Audio.FileID))
		}

		handleAudioUpdate(ctx, msg)
	}

	if update.Message.Photo != nil {
		photo := update.Message.Photo[len(update.Message.Photo)-1]
		path, err := convertImageToPng(downloadFile(ctx, bot, photo.FileID))
		if err == nil {
			msg.AddImage(path)
		}

		handleImageUpdate(ctx, msg)
	}
}
