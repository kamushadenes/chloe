package openai

import (
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/sashabaranov/go-openai"
)

type ChatOptionsOpenAI struct {
	req openai.ChatCompletionRequest
}

func NewChatOptionsOpenAI() schema.ChatOptions {
	return &ChatOptionsOpenAI{req: openai.ChatCompletionRequest{}}
}

func (c *ChatOptionsOpenAI) GetMessages() []schema.Message {
	var msgs []schema.Message
	for k := range c.req.Messages {
		msg := c.req.Messages[k]
		msgs = append(msgs, schema.Message{
			Role:    schema.Role(msg.Role),
			Content: msg.Content,
			Name:    msg.Name,
		})
	}

	return msgs
}

func (c *ChatOptionsOpenAI) GetRequest() openai.ChatCompletionRequest {
	return c.req
}

func (c *ChatOptionsOpenAI) WithMessages(messages []schema.Message) schema.ChatOptions {
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

func (c *ChatOptionsOpenAI) WithModel(model string) schema.ChatOptions {
	c.req.Model = model
	return c
}

func (c *ChatOptionsOpenAI) WithMaxTokens(maxTokens int) schema.ChatOptions {
	c.req.MaxTokens = maxTokens
	return c
}

func (c *ChatOptionsOpenAI) WithTemperature(temperature float32) schema.ChatOptions {
	c.req.Temperature = temperature
	return c
}

func (c *ChatOptionsOpenAI) WithTopP(topP float32) schema.ChatOptions {
	c.req.TopP = topP
	return c
}

func (c *ChatOptionsOpenAI) WithN(n int) schema.ChatOptions {
	c.req.N = n
	return c
}

func (c *ChatOptionsOpenAI) WithStream(stream bool) schema.ChatOptions {
	c.req.Stream = stream
	return c
}

func (c *ChatOptionsOpenAI) WithStop(stop []string) schema.ChatOptions {
	c.req.Stop = stop
	return c
}

func (c *ChatOptionsOpenAI) WithPresencePenalty(presencePenalty float32) schema.ChatOptions {
	c.req.PresencePenalty = presencePenalty
	return c
}

func (c *ChatOptionsOpenAI) WithFrequencyPenalty(frequencyPenalty float32) schema.ChatOptions {
	c.req.FrequencyPenalty = frequencyPenalty
	return c
}

func (c *ChatOptionsOpenAI) WithLogitBias(logitBias map[string]int) schema.ChatOptions {
	c.req.LogitBias = logitBias
	return c
}

func (c *ChatOptionsOpenAI) WithUser(user string) schema.ChatOptions {
	c.req.User = user
	return c
}
