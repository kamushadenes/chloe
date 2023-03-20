package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/messages"
	"github.com/rs/zerolog"
)

func handleTextUpdate(ctx context.Context, msg *messages.Message) {
	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
	_ = processText(ctx, msg, nil)
}

func handleAudioUpdate(ctx context.Context, msg *messages.Message) {
	logger := zerolog.Ctx(ctx)

	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
	tmsg := tgbotapi.NewMessage(msg.Source.Telegram.Update.Message.Chat.ID, "Processing audio...")
	tmsg.ParseMode = tgbotapi.ModeMarkdownV2
	_, _ = msg.Source.Telegram.API.Send(tmsg)

	ch := make(chan interface{})

	if err := processAudio(ctx, msg, ch); err != nil {
		logger.Error().Err(err).Msg("error processing audio")
		return
	}

	transcription := (<-ch).(string)

	oresp := transcription
	response := fmt.Sprintf("Transcription: %s", transcription)

	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
	tmsg = tgbotapi.NewMessage(msg.Source.Telegram.Update.Message.Chat.ID, fmt.Sprintf("Transcription: %s", response))
	tmsg.ParseMode = tgbotapi.ModeMarkdownV2
	_, _ = msg.Source.Telegram.API.Send(tmsg)

	msg.Source.Telegram.Update.Message.Text = oresp

	ch2 := make(chan interface{})
	if err := processText(ctx, msg, ch2); err != nil {
		logger.Error().Err(err).Msg("error processing text")
		return
	}

	response2 := (<-ch2).(string)

	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
	tmsg = tgbotapi.NewMessage(msg.Source.Telegram.Update.Message.Chat.ID, response2)
	tmsg.ParseMode = tgbotapi.ModeMarkdownV2
	_, _ = msg.Source.Telegram.API.Send(tmsg)
}

func handleImageUpdate(ctx context.Context, msg *messages.Message) {
	logger := zerolog.Ctx(ctx)

	tmsg := tgbotapi.NewMessage(msg.Source.Telegram.Update.Message.Chat.ID, "Processing image...")
	tmsg.ParseMode = tgbotapi.ModeMarkdownV2
	_, _ = msg.Source.Telegram.API.Send(tmsg)

	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatUploadPhoto))

	if err := processImage(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error processing image")
	}
}

func handleUpdates(ctx context.Context, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	logger := zerolog.Ctx(ctx)

	l := logger.With().Str("externalUserId", fmt.Sprintf("%d", update.Message.From.ID)).Logger()
	logger = &l

	ctx = logger.WithContext(ctx)

	if update.Message == nil { // ignore non-Message updates
		return
	}

	msg := messages.NewMessage(fmt.Sprintf("%d", update.Message.MessageID))
	msg.Source.Telegram = &messages.TelegramMessageSource{
		API:    bot,
		Update: update,
	}

	user, err := userFromMessage(ctx, msg)
	if err != nil {
		logger.Error().Err(err).Msg("error getting user from message")
		return
	}
	msg.User = user

	channels.IncomingMessagesCh <- msg

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
		for _, photo := range update.Message.Photo {
			msg.AddImage(downloadFile(ctx, bot, photo.FileID))
		}

		handleImageUpdate(ctx, msg)
	}
}
