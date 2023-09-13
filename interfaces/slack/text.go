package slack

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/completion_request_structs"
)

func complete(ctx context.Context, msg *memory.Message) error {
	request := completion_request_structs.NewCompletionRequest()

	request.Message = msg

	request.Context = ctx
	request.Writer = NewSlackWriter(ctx, request, false)

	chat := chat_models.NewChatWithDefaultModel(config.Chat.Provider, msg.User)

	if config.Slack.StreamMessages {
		_, err := chat.ChatStreamWithContext(ctx, request.Writer, messages.UserMessage(promptFromMessage(msg)))
		if err != nil {
			return err
		}
	} else {
		res, err := chat.ChatWithContext(ctx, messages.UserMessage(promptFromMessage(msg)))
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
