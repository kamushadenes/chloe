package openai

import (
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/sashabaranov/go-openai"
	"time"
)

type ChatOptionsOpenAI struct {
	req     openai.ChatCompletionRequest
	timeout time.Duration
}

func NewChatOptionsOpenAI() common.ChatOptions {
	return &ChatOptionsOpenAI{req: openai.ChatCompletionRequest{}}
}

func (c *ChatOptionsOpenAI) GetMessages() []common.Message {
	var msgs []common.Message
	for k := range c.req.Messages {
		msg := c.req.Messages[k]
		msgs = append(msgs, common.Message{
			Role:    common.Role(msg.Role),
			Content: msg.Content,
			Name:    msg.Name,
		})
	}

	return msgs
}

func (c *ChatOptionsOpenAI) GetRequest() interface{} {
	return c.req
}

func (c *ChatOptionsOpenAI) WithMessages(messages []common.Message) common.ChatOptions {
	for k := range messages {
		msg := messages[k]
		c.req.Messages = append(c.req.Messages, openai.ChatCompletionMessage{
			Role:    string(msg.Role),
			Content: msg.Content,
			Name:    msg.Name,
		})
	}

	return c
}

func (c *ChatOptionsOpenAI) WithModel(model string) common.ChatOptions {
	c.req.Model = model
	return c
}

func (c *ChatOptionsOpenAI) WithMaxTokens(maxTokens int) common.ChatOptions {
	c.req.MaxTokens = maxTokens
	return c
}

func (c *ChatOptionsOpenAI) WithTemperature(temperature float32) common.ChatOptions {
	c.req.Temperature = temperature
	return c
}

func (c *ChatOptionsOpenAI) WithTopP(topP float32) common.ChatOptions {
	c.req.TopP = topP
	return c
}

func (c *ChatOptionsOpenAI) WithN(n int) common.ChatOptions {
	c.req.N = n
	return c
}

func (c *ChatOptionsOpenAI) WithStream(stream bool) common.ChatOptions {
	c.req.Stream = stream
	return c
}

func (c *ChatOptionsOpenAI) WithStop(stop []string) common.ChatOptions {
	c.req.Stop = stop
	return c
}

func (c *ChatOptionsOpenAI) WithPresencePenalty(presencePenalty float32) common.ChatOptions {
	c.req.PresencePenalty = presencePenalty
	return c
}

func (c *ChatOptionsOpenAI) WithFrequencyPenalty(frequencyPenalty float32) common.ChatOptions {
	c.req.FrequencyPenalty = frequencyPenalty
	return c
}

func (c *ChatOptionsOpenAI) WithLogitBias(logitBias map[string]int) common.ChatOptions {
	c.req.LogitBias = logitBias
	return c
}

func (c *ChatOptionsOpenAI) WithUser(user string) common.ChatOptions {
	c.req.User = user
	return c
}

func (c *ChatOptionsOpenAI) WithTimeout(dur time.Duration) common.ChatOptions {
	c.timeout = dur
	return c
}

func (c *ChatOptionsOpenAI) GetTimeout() time.Duration {
	return c.timeout
}
