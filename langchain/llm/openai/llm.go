package openai

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/llm/common"
	"github.com/kamushadenes/chloe/logging"
	openai "github.com/sashabaranov/go-openai"
)

type LLMOpenAI struct {
	Client *openai.Client
	Model  *common.LLMModel
}

func NewLLMOpenAI(token string, model *common.LLMModel) common.LLM {
	return &LLMOpenAI{Client: openai.NewClient(token), Model: model}
}

func NewLLMOpenAIWithDefaultModel(token string) common.LLM {
	return NewLLMOpenAI(token, GPT3Dot5TurboInstruct)
}

func (c *LLMOpenAI) Generate(prompt ...string) (common.LLMResult, error) {
	return c.GenerateWithContext(context.Background(), prompt...)
}

func (c *LLMOpenAI) GenerateWithContext(ctx context.Context, prompt ...string) (common.LLMResult, error) {
	opts := NewLLMOptionsOpenAI().
		WithPrompt(prompt).
		WithModel(c.Model.Name).
		WithTimeout(config.Timeouts.Completion)

	return c.GenerateWithOptions(ctx, opts)
}

func (c *LLMOpenAI) GenerateWithOptions(ctx context.Context, opts common.LLMOptions) (common.LLMResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := c.Client.CreateCompletion(ctx, opts.GetRequest().(openai.CompletionRequest))
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

	res.CalculateCosts(c.Model)

	logger.Info().
		Str("provider", "openai").
		Str("model", c.Model.Name).
		Float64("cost", res.Cost.TotalCost).
		Int("tokens", res.Usage.TotalTokens).
		Msg("llm completion done")

	return res, nil
}
