package discord

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
)

func complete(ctx context.Context, msg *memory.Message) error {
	request := structs.NewCompletionRequest()

	request.Message = msg

	request.Context = ctx
	request.Writer = NewDiscordWriter(ctx, request, false)

	chat := chat_models.NewChatWithDefaultModel(config.Chat.Provider, msg.User)

	if config.Discord.StreamMessages {
		_, err := chat.ChatStreamWithContext(ctx, request.Writer, common.UserMessage(promptFromMessage(msg)))
		if err != nil {
			return err
		}
	} else {
		res, err := chat.ChatWithContext(ctx, common.UserMessage(promptFromMessage(msg)))
		if err != nil {
			return err
		}

		for k := range res.Generations {
			_, err = request.Writer.Write([]byte(res.Generations[k].Text))
			if err != nil {
				return err
			}
		}
	}

	return request.Writer.Close()
}
