package openai

import (
	"github.com/kamushadenes/chloe/langchain/llm"
	"github.com/sashabaranov/go-openai"
	"time"
)

type LLMOptionsOpenAI struct {
	req     openai.CompletionRequest
	timeout time.Duration
}

func NewLLMOptionsOpenAI() llm.LLMOptions {
	return &LLMOptionsOpenAI{req: openai.CompletionRequest{}}
}

func (c *LLMOptionsOpenAI) GetPrompt() []string {
	switch p := c.req.Prompt.(type) {
	case string:
		return []string{p}
	case []string:
		return p
	}

	return nil
}

func (c *LLMOptionsOpenAI) GetRequest() openai.CompletionRequest {
	return c.req
}

func (c *LLMOptionsOpenAI) WithPrompt(prompt any) llm.LLMOptions {
	c.req.Prompt = prompt
	return c
}

func (c *LLMOptionsOpenAI) WithModel(model string) llm.LLMOptions {
	c.req.Model = model
	return c
}

func (c *LLMOptionsOpenAI) WithMaxTokens(maxTokens int) llm.LLMOptions {
	c.req.MaxTokens = maxTokens
	return c
}

func (c *LLMOptionsOpenAI) WithTemperature(temperature float32) llm.LLMOptions {
	c.req.Temperature = temperature
	return c
}

func (c *LLMOptionsOpenAI) WithTopP(topP float32) llm.LLMOptions {
	c.req.TopP = topP
	return c
}

func (c *LLMOptionsOpenAI) WithN(n int) llm.LLMOptions {
	c.req.N = n
	return c
}

func (c *LLMOptionsOpenAI) WithStream(stream bool) llm.LLMOptions {
	c.req.Stream = stream
	return c
}

func (c *LLMOptionsOpenAI) WithStop(stop []string) llm.LLMOptions {
	c.req.Stop = stop
	return c
}

func (c *LLMOptionsOpenAI) WithPresencePenalty(presencePenalty float32) llm.LLMOptions {
	c.req.PresencePenalty = presencePenalty
	return c
}

func (c *LLMOptionsOpenAI) WithFrequencyPenalty(frequencyPenalty float32) llm.LLMOptions {
	c.req.FrequencyPenalty = frequencyPenalty
	return c
}

func (c *LLMOptionsOpenAI) WithLogitBias(logitBias map[string]int) llm.LLMOptions {
	c.req.LogitBias = logitBias
	return c
}

func (c *LLMOptionsOpenAI) WithUser(user string) llm.LLMOptions {
	c.req.User = user
	return c
}
func (c *LLMOptionsOpenAI) WithTimeout(dur time.Duration) llm.LLMOptions {
	c.timeout = dur
	return c
}

func (c *LLMOptionsOpenAI) GetTimeout() time.Duration {
	return c.timeout
}
