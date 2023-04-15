package openai

import (
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/sashabaranov/go-openai"
	"time"
)

type ChatOptionsOpenAI struct {
	req     openai.ChatCompletionRequest
	timeout time.Duration
}

func NewChatOptionsOpenAI() chat_models.ChatOptions {
	return &ChatOptionsOpenAI{req: openai.ChatCompletionRequest{}}
}

func (c *ChatOptionsOpenAI) GetMessages() []chat_models.Message {
	var msgs []chat_models.Message
	for k := range c.req.Messages {
		msg := c.req.Messages[k]
		msgs = append(msgs, chat_models.Message{
			Role:    chat_models.Role(msg.Role),
			Content: msg.Content,
			Name:    msg.Name,
		})
	}

	return msgs
}

func (c *ChatOptionsOpenAI) GetRequest() openai.ChatCompletionRequest {
	return c.req
}

func (c *ChatOptionsOpenAI) WithMessages(messages []chat_models.Message) chat_models.ChatOptions {
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

func (c *ChatOptionsOpenAI) WithModel(model string) chat_models.ChatOptions {
	c.req.Model = model
	return c
}

func (c *ChatOptionsOpenAI) WithMaxTokens(maxTokens int) chat_models.ChatOptions {
	c.req.MaxTokens = maxTokens
	return c
}

func (c *ChatOptionsOpenAI) WithTemperature(temperature float32) chat_models.ChatOptions {
	c.req.Temperature = temperature
	return c
}

func (c *ChatOptionsOpenAI) WithTopP(topP float32) chat_models.ChatOptions {
	c.req.TopP = topP
	return c
}

func (c *ChatOptionsOpenAI) WithN(n int) chat_models.ChatOptions {
	c.req.N = n
	return c
}

func (c *ChatOptionsOpenAI) WithStream(stream bool) chat_models.ChatOptions {
	c.req.Stream = stream
	return c
}

func (c *ChatOptionsOpenAI) WithStop(stop []string) chat_models.ChatOptions {
	c.req.Stop = stop
	return c
}

func (c *ChatOptionsOpenAI) WithPresencePenalty(presencePenalty float32) chat_models.ChatOptions {
	c.req.PresencePenalty = presencePenalty
	return c
}

func (c *ChatOptionsOpenAI) WithFrequencyPenalty(frequencyPenalty float32) chat_models.ChatOptions {
	c.req.FrequencyPenalty = frequencyPenalty
	return c
}

func (c *ChatOptionsOpenAI) WithLogitBias(logitBias map[string]int) chat_models.ChatOptions {
	c.req.LogitBias = logitBias
	return c
}

func (c *ChatOptionsOpenAI) WithUser(user string) chat_models.ChatOptions {
	c.req.User = user
	return c
}

func (c *ChatOptionsOpenAI) WithTimeout(dur time.Duration) chat_models.ChatOptions {
	c.timeout = dur
	return c
}

func (c *ChatOptionsOpenAI) GetTimeout() time.Duration {
	return c.timeout
}
