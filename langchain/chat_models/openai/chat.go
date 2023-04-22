package openai

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/logging"
	"github.com/sashabaranov/go-openai"
)

type ChatOpenAI struct {
	client *openai.Client
	model  *common.ChatModel
}

func NewChatOpenAI(token string, model *common.ChatModel) common.Chat {
	return &ChatOpenAI{client: openai.NewClient(token), model: model}
}

func NewChatOpenAIWithDefaultModel(token string) common.Chat {
	return NewChatOpenAI(token, GPT35Turbo)
}

func (c *ChatOpenAI) Chat(messages ...common.Message) (common.ChatResult, error) {
	return c.ChatWithContext(context.Background(), messages...)
}

func (c *ChatOpenAI) ChatWithContext(ctx context.Context, messages ...common.Message) (common.ChatResult, error) {
	opts := NewChatOptionsOpenAI().WithMessages(messages).WithModel(c.model.Name)

	return c.ChatWithOptions(ctx, opts)
}

func (c *ChatOpenAI) ChatWithOptions(ctx context.Context, opts common.ChatOptions) (common.ChatResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := c.client.CreateChatCompletion(ctx, opts.GetRequest())
	if err != nil {
		return common.ChatResult{}, err
	}

	var res common.ChatResult

	for k := range resp.Choices {
		res.Generations = append(res.Generations, common.ChatGeneration{
			Text: resp.Choices[k].Message.Content,
			Message: common.Message{
				Name:    resp.Choices[k].Message.Name,
				Role:    common.Role(resp.Choices[k].Message.Role),
				Content: resp.Choices[k].Message.Content,
			},
		})
	}

	res.Usage = common.ChatUsage{
		PromptTokens:     resp.Usage.PromptTokens,
		CompletionTokens: resp.Usage.CompletionTokens,
		TotalTokens:      resp.Usage.TotalTokens,
	}

	res.CalculateCosts(c.model)

	logger.Info().
		Str("provider", "openai").
		Str("model", c.model.Name).
		Float64("cost", res.Cost.TotalCost).
		Int("tokens", res.Usage.TotalTokens).
		Msg("chat completion done")

	return res, nil
}
