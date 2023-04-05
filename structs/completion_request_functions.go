package structs

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/resources"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"time"
)

func (creq *CompletionRequest) GetID() string {
	return creq.ID
}

func (creq *CompletionRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *CompletionRequest) Copy() *CompletionRequest {
	return &CompletionRequest{
		Context:         creq.Context,
		Writer:          creq.Writer,
		SkipClose:       creq.SkipClose,
		StartChannel:    creq.StartChannel,
		ContinueChannel: creq.ContinueChannel,
		ErrorChannel:    creq.ErrorChannel,
		ResultChannel:   creq.ResultChannel,
		Message:         creq.Message.Copy(),
		Mode:            creq.Mode,
		Args:            creq.Args,
	}
}

func (creq *CompletionRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *CompletionRequest) GetWriter() ChloeWriter {
	return creq.Writer
}

func (creq *CompletionRequest) SetWriter(w ChloeWriter) {
	creq.Writer = w
}

func (creq *CompletionRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *CompletionRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *CompletionRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *CompletionRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *CompletionRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}

func (creq *CompletionRequest) GetCost() float64 {
	return config.OpenAI.GetModel(config.Completion).GetCostForTokens(creq.CountTokens())
}

// CountTokens based on https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb
func (creq *CompletionRequest) CountTokens() int {
	messages := creq.ToChatCompletionMessages()

	return creq.CountChatCompletionTokens(messages)
}

func (creq *CompletionRequest) CountChatCompletionTokens(messages []openai.ChatCompletionMessage) int {
	model := config.OpenAI.GetModel(config.Completion)

	return model.CountChatCompletionTokens(messages)
}

func (creq *CompletionRequest) getArgs() map[string]interface{} {
	args := creq.Args
	if args == nil {
		args = make(map[string]interface{})
	}
	args["Interface"] = creq.Message.Interface
	args["User"] = creq.Message.User

	args["Date"] = time.Now().Format("2006-01-02")
	args["Time"] = time.Now().Format("15:04:05")

	return args
}

func (creq *CompletionRequest) getSystemMessages() []openai.ChatCompletionMessage {
	logger := zerolog.Ctx(creq.GetContext())

	var messages []openai.ChatCompletionMessage

	args := creq.getArgs()

	// Load bootstrap values
	bootstrap, err := resources.GetPrompt("bootstrap", &resources.PromptArgs{
		Args: args,
		Mode: creq.Mode,
	})
	if err != nil {
		logger.Error().Err(err).Msg("failed to load bootstrap prompt")
	}
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "system",
		Content: bootstrap,
	})

	// Load system prompt
	prompt, err := resources.GetPrompt(creq.Mode, &resources.PromptArgs{Args: args, Mode: creq.Mode})
	if err != nil {
		panic(err)
	}
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "system",
		Content: prompt,
	})

	// Feed few-shot examples, if any
	examples, err := resources.GetExamples(prompt, &resources.PromptArgs{
		Args: args,
		Mode: creq.Mode,
	})
	if err == nil {
		messages = append(messages, examples...)
	}

	return messages
}

func (creq *CompletionRequest) getUserMessages() []openai.ChatCompletionMessage {
	logger := zerolog.Ctx(creq.GetContext())

	var messages []openai.ChatCompletionMessage

	savedMessages, err := creq.Message.User.ListMessages(creq.GetContext())
	if err != nil {
		logger.Error().Err(err).Msg("failed to load saved messages")
		return nil
	}

	for k, m := range savedMessages {
		var content string

		if k >= len(savedMessages)-config.OpenAI.MessagesToKeepFullContent {
			content = m.Content
		} else {
			content = m.GetContent()
		}

		if len(content) > 0 {
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    m.Role,
				Content: content,
			})
		}
	}

	return messages
}

func (creq *CompletionRequest) ToChatCompletionMessages() []openai.ChatCompletionMessage {
	systemMessages := creq.getSystemMessages()
	userMessages := creq.getUserMessages()

	systemCount := creq.CountChatCompletionTokens(systemMessages)
	userCount := creq.CountChatCompletionTokens(userMessages)

	for {
		if (systemCount + userCount) > config.OpenAI.GetModel(config.Completion).GetContextSize() {
			userMessages = userMessages[1:]
			userCount = creq.CountChatCompletionTokens(userMessages)
		} else {
			break
		}
	}

	return append(systemMessages, userMessages...)
}
