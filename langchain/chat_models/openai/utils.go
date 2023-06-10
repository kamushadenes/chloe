package openai

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/tokenizer"
)

func (c *ChatOpenAI) LoadUserMessages(ctx context.Context) ([]common.Message, error) {
	var messages []common.Message

	msgs, err := c.User.ListMessages(ctx)
	if err != nil {
		return nil, err
	}

	for k := range msgs {
		m := common.Message{
			Role: common.Role(msgs[k].Role),
		}

		if k >= len(messages)-config.OpenAI.MessagesToKeepFullContent {
			m.Content = msgs[k].Content
		} else {
			m.Content = msgs[k].GetContent()
		}

		messages = append(messages, m)
	}

	return messages, nil
}

func (c *ChatOpenAI) ReduceTokens(systemMessages []common.Message, messages []common.Message) []common.Message {
	modelName := c.Model.Tokenizer
	if c.Model.Tokenizer == "" {
		modelName = c.Model.Name
	}

	for {
		var tokenCount int
		for k := range systemMessages {
			tokenCount += tokenizer.CountTokens(modelName, systemMessages[k].Content)
		}

		for k := range messages {
			tokenCount += tokenizer.CountTokens(modelName, messages[k].Content)
		}

		if tokenCount > c.Model.ContextSize {
			messages = messages[1:]
			continue
		}

		return append(systemMessages, messages...)
	}
}
