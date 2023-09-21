package slack

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/chat_models/base"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func complete(ctx context.Context, msg *memory.Message) error {
	w := NewSlackWriter(ctx, msg, false)

	chat := base.NewChatWithDefaultModel(config.Chat.Provider, msg.User)

	if config.Slack.StreamMessages {
		if _, err := chat.ChatStreamWithContext(ctx, w, msg, messages.UserMessage(promptFromMessage(msg))); err != nil {
			return err
		}
	} else {
		res, err := chat.ChatWithContext(ctx, msg, messages.UserMessage(promptFromMessage(msg)))
		if err != nil {
			return err
		}

		for k := range res.Generations {
			if _, err = w.Write([]byte(res.Generations[k].Text)); err != nil {
				return err
			}
		}
	}

	return w.Close()
}
