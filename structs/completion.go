package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/resources"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
	"time"
)

type CompletionRequest struct {
	ID      string
	Context context.Context

	Writer    io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message *memory.Message

	Mode string                 `json:"mode"`
	Args map[string]interface{} `json:"args"`
}

func NewCompletionRequest() *CompletionRequest {
	return &CompletionRequest{
		ID:   uuid.Must(uuid.NewV4()).String(),
		Args: make(map[string]interface{}),
	}
}

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

func (creq *CompletionRequest) GetWriters() []io.WriteCloser {
	return []io.WriteCloser{creq.Writer}
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

func (creq *CompletionRequest) CountTokens(messages []openai.ChatCompletionMessage) int {
	var tokens int
	tokens += 2 // every reply is primed with <im_start>assistant
	for k := range messages {
		tokens += 4 // every message follows <im_start>{role/name}\n{content}<im_end>\n
		if messages[k].Name != "" && messages[k].Role == "" {
			tokens -= 1 // if there's a name, the role can be ommited, so we need to remove one token if it's empty
		}
		tokens += int(float64(len(strings.Fields(messages[k].Content))) * 0.75) // count words
	}

	return tokens
}

func (creq *CompletionRequest) ToChatCompletionMessages() []openai.ChatCompletionMessage {
	logger := zerolog.Ctx(creq.GetContext())

	var messages []openai.ChatCompletionMessage

	args := creq.Args
	if args == nil {
		args = make(map[string]interface{})
	}
	args["Interface"] = creq.Message.Interface
	args["User"] = creq.Message.User

	args["Date"] = time.Now().Format("2006-01-02")
	args["Time"] = time.Now().Format("15:04:05")

	prompt, err := resources.GetPrompt(creq.Mode, &resources.PromptArgs{Args: args, Mode: creq.Mode})
	if err != nil {
		panic(err)
	}

	// messages = append(messages, openai.ChatCompletionMessage{Role: "system", Content: prompt})
	messages = append(messages, openai.ChatCompletionMessage{Role: "user", Content: prompt})

	bootstrap, err := resources.GetPrompt("bootstrap", &resources.PromptArgs{Args: args, Mode: creq.Mode})
	if err != nil {
		logger.Error().Err(err).Msg("failed to load bootstrap prompt")
	}
	// messages = append(messages, openai.ChatCompletionMessage{Role: "system", Content: bootstrap})
	messages = append(messages, openai.ChatCompletionMessage{Role: "user", Content: bootstrap})

	var userMessages []openai.ChatCompletionMessage

	savedMessages, err := creq.Message.User.LoadMessages(creq.GetContext())
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
			userMessages = append(userMessages, openai.ChatCompletionMessage{Role: m.Role, Content: content})
		}
	}

	systemCount := creq.CountTokens(messages)
	userCount := creq.CountTokens(userMessages)

	for {
		if (systemCount + userCount) > config.OpenAI.MaxTokens[config.OpenAI.DefaultModel.Completion] {
			userMessages = userMessages[1:]
			userCount = creq.CountTokens(userMessages)
		} else {
			break
		}
	}

	messages = append(messages, userMessages...)

	return messages
}
