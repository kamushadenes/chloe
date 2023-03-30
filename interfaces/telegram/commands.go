package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/resources"
	"github.com/rs/zerolog"
	"strings"
)

func handleCommands(ctx context.Context, msg *memory.Message) bool {
	switch msg.Source.Telegram.Update.Message.Command() {
	case "start":
		_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
		tryAndRespond(ctx, msg, "Hi!", "", nil, true)
		return true
	case "mode":
		_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
		err := setUserMode(ctx, msg)
		tryAndRespond(ctx, msg, "Mode set!", "Failed to set mode", err, true)
		return true
	case "forget":
		_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
		err := forgetUser(ctx, msg)
		tryAndRespond(ctx, msg, i18n.GetForgetText(), "Failed to forget user", err, true)
		return true
	case "listmodes":
		_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
		modes, err := listModes()
		tryAndRespond(ctx, msg, modes, "Failed to list modes", err, true)
		return true
	case "generate":
		if msg.Source.Telegram.Update.Message.CommandArguments() == "" {
			tryAndRespond(ctx, msg, "", "Please provide a prompt!",
				fmt.Errorf("missing arguments"), true)
			return true
		}

		generate(ctx, msg)
		return true
	case "tts":
		if msg.Source.Telegram.Update.Message.CommandArguments() == "" {
			tryAndRespond(ctx, msg, "", "Please provide a prompt!",
				fmt.Errorf("missing arguments"), true)
			return true
		}

		tts(ctx, msg)
		return true
	}

	return false
}

func action(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatUploadPhoto))

	if err := aiAction(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error running action")
	}
}
func generate(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatUploadPhoto))

	if err := aiGenerate(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating image")
	}
}

func tts(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatRecordVoice))

	if err := aiTTS(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating audio")
	}
}

func forgetUser(ctx context.Context, msg *memory.Message) error {
	return msg.User.DeleteMessages(ctx)
}

func setUserMode(ctx context.Context, msg *memory.Message) error {
	return msg.User.SetMode(ctx, msg.Source.Telegram.Update.Message.CommandArguments())
}

func listModes() (string, error) {
	modes, err := resources.ListPrompts()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Available modes:\n\n%s", strings.Join(modes, "\n")), nil
}
