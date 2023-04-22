package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/chat_models/openai"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
)

func aiComplete(ctx context.Context, msg *memory.Message, ch chan interface{}) error {
	request := structs.NewCompletionRequest()

	request.Message = msg

	request.ResultChannel = ch
	request.Context = ctx
	request.Writer = NewTelegramWriter(ctx, request, false)

	chat := openai.NewChatOpenAIWithDefaultModel(config.OpenAI.APIKey)

	if config.Slack.StreamMessages {
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
