package openai

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/sashabaranov/go-openai"
)

type ChatOpenAI struct {
	client *openai.Client
	model  *chat_models.ChatModel
}

func NewChatOpenAI(token string, model *chat_models.ChatModel) chat_models.Chat {
	return &ChatOpenAI{client: openai.NewClient(token), model: model}
}

func NewChatOpenAIWithDefaultModel(token string) chat_models.Chat {
	return NewChatOpenAI(token, GPT35Turbo)
}

func (c *ChatOpenAI) Chat(messages ...chat_models.Message) (chat_models.ChatResult, error) {
	return c.ChatWithContext(context.Background(), messages...)
}

func (c *ChatOpenAI) ChatWithContext(ctx context.Context, messages ...chat_models.Message) (chat_models.ChatResult, error) {
	opts := NewChatOptionsOpenAI().WithMessages(messages).WithModel(c.model.Name)

	return c.ChatWithOptions(ctx, opts)
}

func (c *ChatOpenAI) ChatWithOptions(ctx context.Context, opts chat_models.ChatOptions) (chat_models.ChatResult, error) {
	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := c.client.CreateChatCompletion(ctx, opts.GetRequest())
	if err != nil {
		return chat_models.ChatResult{}, err
	}

	var res chat_models.ChatResult

	for k := range resp.Choices {
		res.Generations = append(res.Generations, chat_models.ChatGeneration{
			Text: resp.Choices[k].Message.Content,
			Message: chat_models.Message{
				Name:    resp.Choices[k].Message.Name,
				Role:    chat_models.Role(resp.Choices[k].Message.Role),
				Content: resp.Choices[k].Message.Content,
			},
		})
	}

	res.Usage = chat_models.ChatUsage{
		PromptTokens:     resp.Usage.PromptTokens,
		CompletionTokens: resp.Usage.CompletionTokens,
		TotalTokens:      resp.Usage.TotalTokens,
	}

	res.CalculateCosts(c.model)

	return res, nil
}
