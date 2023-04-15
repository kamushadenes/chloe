package openai

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/llm"
	"github.com/sashabaranov/go-openai"
)

type LLMOpenAI struct {
	client *openai.Client
	model  *llm.LLMModel
}

func NewLLMOpenAI(token string, model *llm.LLMModel) llm.LLM {
	return &LLMOpenAI{client: openai.NewClient(token), model: model}
}

func NewLLMOpenAIWithDefaultModel(token string) llm.LLM {
	return NewLLMOpenAI(token, TextDavinci003)
}

func (c *LLMOpenAI) Generate(prompt ...string) (llm.LLMResult, error) {
	return c.GenerateWithContext(context.Background(), prompt...)
}

func (c *LLMOpenAI) GenerateWithContext(ctx context.Context, prompt ...string) (llm.LLMResult, error) {
	opts := NewLLMOptionsOpenAI().WithPrompt(prompt).WithModel(c.model.Name)

	return c.GenerateWithOptions(ctx, opts)
}

func (c *LLMOpenAI) GenerateWithOptions(ctx context.Context, opts llm.LLMOptions) (llm.LLMResult, error) {
	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := c.client.CreateCompletion(ctx, opts.GetRequest())
	if err != nil {
		return llm.LLMResult{}, err
	}

	var res llm.LLMResult

	for k := range resp.Choices {
		res.Generations = append(res.Generations, llm.LLMGeneration{
			Text:         resp.Choices[k].Text,
			Index:        resp.Choices[k].Index,
			FinishReason: resp.Choices[k].FinishReason,
		})
	}

	res.Usage = llm.LLMUsage{
		PromptTokens:     resp.Usage.PromptTokens,
		CompletionTokens: resp.Usage.CompletionTokens,
		TotalTokens:      resp.Usage.TotalTokens,
	}

	return res, nil
}
