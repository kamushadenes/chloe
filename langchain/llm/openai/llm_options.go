package openai

import (
	"github.com/kamushadenes/chloe/langchain/llm/common"
	"github.com/sashabaranov/go-openai"
	"time"
)

type LLMOptionsOpenAI struct {
	req     openai.CompletionRequest
	timeout time.Duration
}

func NewLLMOptionsOpenAI() common.LLMOptions {
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

func (c *LLMOptionsOpenAI) WithPrompt(prompt any) common.LLMOptions {
	c.req.Prompt = prompt
	return c
}

func (c *LLMOptionsOpenAI) WithModel(model string) common.LLMOptions {
	c.req.Model = model
	return c
}

func (c *LLMOptionsOpenAI) WithMaxTokens(maxTokens int) common.LLMOptions {
	c.req.MaxTokens = maxTokens
	return c
}

func (c *LLMOptionsOpenAI) WithTemperature(temperature float32) common.LLMOptions {
	c.req.Temperature = temperature
	return c
}

func (c *LLMOptionsOpenAI) WithTopP(topP float32) common.LLMOptions {
	c.req.TopP = topP
	return c
}

func (c *LLMOptionsOpenAI) WithN(n int) common.LLMOptions {
	c.req.N = n
	return c
}

func (c *LLMOptionsOpenAI) WithStream(stream bool) common.LLMOptions {
	c.req.Stream = stream
	return c
}

func (c *LLMOptionsOpenAI) WithStop(stop []string) common.LLMOptions {
	c.req.Stop = stop
	return c
}

func (c *LLMOptionsOpenAI) WithPresencePenalty(presencePenalty float32) common.LLMOptions {
	c.req.PresencePenalty = presencePenalty
	return c
}

func (c *LLMOptionsOpenAI) WithFrequencyPenalty(frequencyPenalty float32) common.LLMOptions {
	c.req.FrequencyPenalty = frequencyPenalty
	return c
}

func (c *LLMOptionsOpenAI) WithLogitBias(logitBias map[string]int) common.LLMOptions {
	c.req.LogitBias = logitBias
	return c
}

func (c *LLMOptionsOpenAI) WithUser(user string) common.LLMOptions {
	c.req.User = user
	return c
}
func (c *LLMOptionsOpenAI) WithTimeout(dur time.Duration) common.LLMOptions {
	c.timeout = dur
	return c
}

func (c *LLMOptionsOpenAI) GetTimeout() time.Duration {
	return c.timeout
}
