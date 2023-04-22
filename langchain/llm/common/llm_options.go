package common

import (
	"time"
)

type LLMOptions interface {
	GetRequest() interface{}
	WithPrompt(any) LLMOptions
	WithModel(model string) LLMOptions
	WithMaxTokens(maxTokens int) LLMOptions
	WithTemperature(temperature float32) LLMOptions
	WithTopP(topP float32) LLMOptions
	WithN(n int) LLMOptions
	WithStream(stream bool) LLMOptions
	WithStop(stop []string) LLMOptions
	WithPresencePenalty(presencePenalty float32) LLMOptions
	WithFrequencyPenalty(frequencyPenalty float32) LLMOptions
	WithLogitBias(logitBias map[string]int) LLMOptions
	WithUser(user string) LLMOptions
	GetPrompt() []string
	WithTimeout(time.Duration) LLMOptions
	GetTimeout() time.Duration
}
