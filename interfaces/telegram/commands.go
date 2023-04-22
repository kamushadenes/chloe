package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func handleCommands(ctx context.Context, msg *memory.Message) bool {
	switch msg.Source.Telegram.Update.Message.Command() {
	case "start":
		_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
		tryAndRespond(ctx, msg, "Hi!", "", nil, true)
		return true
	case "forget":
		_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatTyping))
		err := forgetUser(ctx, msg)
		tryAndRespond(ctx, msg, i18n.GetForgetText(), "Failed to forget user", err, true)
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
	case "action":
		if msg.Source.Telegram.Update.Message.CommandArguments() == "" {
			tryAndRespond(ctx, msg, "", "Please provide a prompt!",
				fmt.Errorf("missing arguments"), true)
			return true
		}

		action(ctx, msg)
		return true
	}

	return false
}

func action(ctx context.Context, msg *memory.Message) {
	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatUploadPhoto))

	if err := aiAction(ctx, msg); err != nil {
		_ = msg.SendError(err)
	}
}
func generate(ctx context.Context, msg *memory.Message) {
	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatUploadPhoto))

	if err := aiGenerate(ctx, msg); err != nil {
		_ = msg.SendError(err)
	}
}

func tts(ctx context.Context, msg *memory.Message) {
	_, _ = msg.Source.Telegram.API.Send(tgbotapi.NewChatAction(msg.Source.Telegram.Update.Message.Chat.ID, tgbotapi.ChatRecordVoice))

	if err := aiTTS(ctx, msg); err != nil {
		_ = msg.SendError(err)
	}
}

func forgetUser(ctx context.Context, msg *memory.Message) error {
	if err := msg.User.DeleteMessages(ctx); err != nil {
		_ = msg.SendError(err)
		return err
	}

	return nil
}
