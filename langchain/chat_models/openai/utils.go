package openai

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/tokenizer"
)

func (c *ChatOpenAI) LoadUserMessages(ctx context.Context) ([]messages.Message, error) {
	var msgs []messages.Message

	mmsgs, err := c.User.ListMessages(ctx)
	if err != nil {
		return nil, err
	}

	for k := range mmsgs {
		m := messages.Message{
			Role: messages.Role(mmsgs[k].Role),
		}

		if k >= len(mmsgs)-config.OpenAI.MessagesToKeepFullContent {
			m.Content = mmsgs[k].Content
		} else {
			m.Content = mmsgs[k].GetContent()
		}

		msgs = append(msgs, m)
	}

	return msgs, nil
}

func (c *ChatOpenAI) ReduceTokens(systemMessages []messages.Message, msgs []messages.Message) []messages.Message {
	modelName := c.Model.Tokenizer
	if c.Model.Tokenizer == "" {
		modelName = c.Model.Name
	}

	for {
		var tokenCount int
		for k := range systemMessages {
			tokenCount += tokenizer.CountTokens(modelName, systemMessages[k].Content)
		}

		for k := range msgs {
			tokenCount += tokenizer.CountTokens(modelName, msgs[k].Content)
		}

		if tokenCount > c.Model.ContextSize {
			msgs = msgs[1:]
			continue
		}

		return append(systemMessages, msgs...)
	}
}
