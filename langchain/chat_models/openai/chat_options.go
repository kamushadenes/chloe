package openai

import (
	"time"

	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/prompts"
	"github.com/sashabaranov/go-openai"
)

type ChatOptionsOpenAI struct {
	req       openai.ChatCompletionRequest
	timeout   time.Duration
	bootstrap messages.Message
	system    messages.Message
	examples  []messages.Message
}

func NewChatOptionsOpenAI() common.ChatOptions {
	opts := &ChatOptionsOpenAI{req: openai.ChatCompletionRequest{}}

	return opts.
		WithBootstrap(prompts.BootstrapArgs{
			Interface:     "unknown",
			UserID:        0,
			UserFirstName: "User",
			UserLastName:  "",
		}).
		WithSystemPrompt("default").
		WithExamples("default")
}

func (c ChatOptionsOpenAI) GetMessages() []messages.Message {
	var msgs []messages.Message
	for k := range c.req.Messages {
		msg := c.req.Messages[k]
		msgs = append(msgs, messages.Message{
			Role:    messages.Role(msg.Role),
			Content: msg.Content,
			Name:    msg.Name,
		})
	}

	return msgs
}

func (c ChatOptionsOpenAI) GetSystemMessages() []messages.Message {
	var msgs []messages.Message
	msgs = append(msgs, c.bootstrap)
	msgs = append(msgs, c.system)
	msgs = append(msgs, c.examples...)

	return msgs
}

func (c ChatOptionsOpenAI) GetRequest() interface{} {
	return c.req
}

func (c ChatOptionsOpenAI) WithMessages(messages []messages.Message) common.ChatOptions {
	var msgs []openai.ChatCompletionMessage

	for k := range messages {
		msg := messages[k]
		msgs = append(msgs, openai.ChatCompletionMessage{
			Role:    string(msg.Role),
			Content: msg.Content,
			Name:    msg.Name,
		})
	}

	c.req.Messages = msgs

	return c
}

func (c ChatOptionsOpenAI) WithModel(model string) common.ChatOptions {
	c.req.Model = model
	return c
}

func (c ChatOptionsOpenAI) WithMaxTokens(maxTokens int) common.ChatOptions {
	c.req.MaxTokens = maxTokens
	return c
}

func (c ChatOptionsOpenAI) WithTemperature(temperature float32) common.ChatOptions {
	c.req.Temperature = temperature
	return c
}

func (c ChatOptionsOpenAI) WithTopP(topP float32) common.ChatOptions {
	c.req.TopP = topP
	return c
}

func (c ChatOptionsOpenAI) WithN(n int) common.ChatOptions {
	c.req.N = n
	return c
}

func (c ChatOptionsOpenAI) WithStream(stream bool) common.ChatOptions {
	c.req.Stream = stream
	return c
}

func (c ChatOptionsOpenAI) WithStop(stop []string) common.ChatOptions {
	c.req.Stop = stop
	return c
}

func (c ChatOptionsOpenAI) WithPresencePenalty(presencePenalty float32) common.ChatOptions {
	c.req.PresencePenalty = presencePenalty
	return c
}

func (c ChatOptionsOpenAI) WithFrequencyPenalty(frequencyPenalty float32) common.ChatOptions {
	c.req.FrequencyPenalty = frequencyPenalty
	return c
}

func (c ChatOptionsOpenAI) WithLogitBias(logitBias map[string]int) common.ChatOptions {
	c.req.LogitBias = logitBias
	return c
}

func (c ChatOptionsOpenAI) WithUser(user string) common.ChatOptions {
	c.req.User = user
	return c
}

func (c ChatOptionsOpenAI) WithTimeout(dur time.Duration) common.ChatOptions {
	c.timeout = dur
	return c
}

func (c ChatOptionsOpenAI) GetTimeout() time.Duration {
	return c.timeout
}

func (c ChatOptionsOpenAI) WithSystemPrompt(promptName string) common.ChatOptions {
	if prompt, err := prompts.GetPrompt(promptName, make(map[string]interface{})); err == nil {
		c.system = messages.SystemMessage(prompt)
	}

	return c
}

func (c ChatOptionsOpenAI) WithBootstrap(args interface{}) common.ChatOptions {
	switch v := args.(type) {
	case prompts.BootstrapArgs:
		if v.Date == "" {
			v.Date = time.Now().Format("2006-01-02")
		}

		if v.Time == "" {
			v.Time = time.Now().Format("15:04:05")
		}

		if bootstrap, err := prompts.GetPrompt("bootstrap", v); err == nil {
			c.bootstrap = messages.SystemMessage(bootstrap)
		}
	default:
		if bootstrap, err := prompts.GetPrompt("bootstrap", v); err == nil {
			c.bootstrap = messages.SystemMessage(bootstrap)
		}
	}

	return c
}

func (c ChatOptionsOpenAI) WithExamples(promptName string) common.ChatOptions {
	/*
		if examples, err := prompts.GetExamples(promptName, make(map[string]interface{})); err == nil {
			c.examples = examples
		}
	*/

	return c
}

func (c ChatOptionsOpenAI) WithFunctions(functions []*functions.FunctionDefinition) common.ChatOptions {
	for k := range functions {
		c.req.Functions = append(c.req.Functions, openai.FunctionDefinition{
			Name:        functions[k].Name,
			Description: functions[k].Description,
			Parameters:  functions[k].Parameters,
		})
	}

	return c
}

func (c ChatOptionsOpenAI) GetFunctions() []*functions.FunctionDefinition {
	var fns []*functions.FunctionDefinition

	for k := range c.req.Functions {
		fns = append(fns, &functions.FunctionDefinition{
			Name:        c.req.Functions[k].Name,
			Description: c.req.Functions[k].Description,
			Parameters:  c.req.Functions[k].Parameters,
		})
	}

	return fns
}
