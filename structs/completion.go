package structs

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/resources"
	"github.com/kamushadenes/chloe/users"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"io"
	"time"
)

type CompletionRequest struct {
	Context context.Context

	Writer    io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	User    *users.User            `json:"user,omitempty"`
	Content string                 `json:"content"`
	Summary string                 `json:"summary"`
	Mode    string                 `json:"mode"`
	Args    map[string]interface{} `json:"args"`
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
		User:            creq.User,
		Content:         creq.Content,
		Summary:         creq.Summary,
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

func (creq *CompletionRequest) GetTokenCount(chainOfThought bool) int {
	messages := creq.ToChatCompletionMessages(creq.Context, chainOfThought)

	var tokens int
	for k := range messages {
		tokens += int(float64(len(messages[k].Content)) * 0.75)
	}

	return tokens
}

func (creq *CompletionRequest) ToChatCompletionMessages(ctx context.Context, chainOfThought bool) []openai.ChatCompletionMessage {
	logger := zerolog.Ctx(ctx)

	var messages []openai.ChatCompletionMessage

	args := creq.Args
	if args == nil {
		args = make(map[string]interface{})
	}
	args["User"] = creq.User

	args["Date"] = time.Now().Format("2006-01-02")
	args["Time"] = time.Now().Format("15:04:05")

	prompt, err := resources.GetPrompt(creq.Mode, &resources.PromptArgs{Args: args, Mode: creq.Mode})
	if err != nil {
		panic(err)
	}

	messages = append(messages, openai.ChatCompletionMessage{Role: "system", Content: prompt})

	bootstrap, err := resources.GetPrompt("bootstrap", &resources.PromptArgs{Args: args, Mode: creq.Mode})
	if err != nil {
		logger.Error().Err(err).Msg("failed to load bootstrap prompt")
	}
	messages = append(messages, openai.ChatCompletionMessage{Role: "system", Content: bootstrap})

	var userMessages []openai.ChatCompletionMessage

	savedMessages, err := memory.LoadMessages(ctx, creq.User.ID)
	if err != nil {
		logger.Error().Err(err).Msg("failed to load saved messages")
		return nil
	}
	var tokens float64
	for k := range messages {
		tokens += float64(len(messages[k].Content)) * 0.75
	}
	tokens += float64(len(creq.Content)) * 0.75

	for k, m := range savedMessages {
		role := m[1]
		var content string

		if chainOfThought {
			if len(m[4]) > 0 {
				content = m[4]
			}
		}
		if len(content) == 0 {
			if k >= len(savedMessages)-4 || (len(m[2]) < len(m[3]) || len(m[3]) == 0) {
				content = m[2]
			} else {
				content = m[3]
			}
		}

		if len(content) > 0 {
			userMessages = append(userMessages, openai.ChatCompletionMessage{Role: role, Content: content})
		}
		tokens += float64(len(content)) * 0.75
	}

	for {
		if tokens > 4096 {
			tokens -= float64(len(userMessages[0].Content)) * 0.75
			userMessages = userMessages[1:]
		} else {
			break
		}
	}

	messages = append(messages, userMessages...)

	err = memory.SaveMessage(ctx, creq.User.ID, "user", creq.Content, "")
	if err != nil {
		panic(err)
	}

	messages = append(messages, openai.ChatCompletionMessage{Role: "user", Content: creq.Content})

	return messages
}
