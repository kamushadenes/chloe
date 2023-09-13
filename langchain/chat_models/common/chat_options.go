package common

import (
	"time"

	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
)

type ChatOptions interface {
	GetRequest() interface{}
	WithMessages(messages []messages.Message) ChatOptions
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
	GetSystemMessages() []messages.Message
	GetMessages() []messages.Message
	WithTimeout(time.Duration) ChatOptions
	GetTimeout() time.Duration
	WithSystemPrompt(promptName string) ChatOptions
	WithBootstrap(args interface{}) ChatOptions
	WithExamples(promptName string) ChatOptions
	WithFunctions([]*functions.FunctionDefinition) ChatOptions
	GetFunctions() []*functions.FunctionDefinition
	WithUserMessage(message *memory.Message) ChatOptions
	GetUserMessage() *memory.Message
}
