package chat_models

import (
	"github.com/sashabaranov/go-openai"
	"time"
)

type ChatOptions interface {
	GetRequest() openai.ChatCompletionRequest
	WithMessages(messages []Message) ChatOptions
	WithModel(model string) ChatOptions
	WithMaxTokens(maxTokens int) ChatOptions
	WithTemperature(temperature float32) ChatOptions
	WithTopP(topP float32) ChatOptions
	WithN(n int) ChatOptions
	WithStream(stream bool) ChatOptions
	WithStop(stop []string) ChatOptions
	WithPresencePenalty(presencePenalty float32) ChatOptions
	WithFrequencyPenalty(frequencyPenalty float32) ChatOptions
	WithLogitBias(logitBias map[string]int) ChatOptions
	WithUser(user string) ChatOptions
	GetMessages() []Message
	WithTimeout(time.Duration) ChatOptions
	GetTimeout() time.Duration
}