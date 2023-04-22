package openai

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/llm/common"
	"github.com/kamushadenes/chloe/logging"
	"github.com/sashabaranov/go-openai"
)

type LLMOpenAI struct {
	client *openai.Client
	model  *common.LLMModel
}

func NewLLMOpenAI(token string, model *common.LLMModel) common.LLM {
	return &LLMOpenAI{client: openai.NewClient(token), model: model}
}

func NewLLMOpenAIWithDefaultModel(token string) common.LLM {
	return NewLLMOpenAI(token, TextDavinci003)
}

func (c *LLMOpenAI) Generate(prompt ...string) (common.LLMResult, error) {
	return c.GenerateWithContext(context.Background(), prompt...)
}

func (c *LLMOpenAI) GenerateWithContext(ctx context.Context, prompt ...string) (common.LLMResult, error) {
	opts := NewLLMOptionsOpenAI().WithPrompt(prompt).WithModel(c.model.Name)

	return c.GenerateWithOptions(ctx, opts)
}

func (c *LLMOpenAI) GenerateWithOptions(ctx context.Context, opts common.LLMOptions) (common.LLMResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := c.client.CreateCompletion(ctx, opts.GetRequest())
	if err != nil {
		return common.LLMResult{}, err
	}

	var res common.LLMResult

	for k := range resp.Choices {
		res.Generations = append(res.Generations, common.LLMGeneration{
			Text:         resp.Choices[k].Text,
			Index:        resp.Choices[k].Index,
			FinishReason: resp.Choices[k].FinishReason,
		})
	}

	res.Usage = common.LLMUsage{
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
		Msg("llm completion done")

	return res, nil
}
