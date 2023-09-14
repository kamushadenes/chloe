package telegram

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func aiComplete(ctx context.Context, msg *memory.Message) error {
	w := NewTelegramWriter(ctx, msg, false)

	chat := chat_models.NewChatWithDefaultModel(config.Chat.Provider, msg.User)

	if config.Telegram.StreamMessages {
		_, err := chat.ChatStreamWithContext(ctx, w, msg, messages.UserMessage(promptFromMessage(msg)))
		if err != nil {
			return err
		}
	} else {
		res, err := chat.ChatWithContext(ctx, msg, messages.UserMessage(promptFromMessage(msg)))
		if err != nil {
			return err
		}

		for k := range res.Generations {
			_, err = w.Write([]byte(res.Generations[k].Text))
			if err != nil {
				return err
			}
		}
	}

	return w.Close()
}
