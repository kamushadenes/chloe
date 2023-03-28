package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
	"github.com/sashabaranov/go-openai"
	"io"
)

type ActionRequest struct {
	ID      string
	Context context.Context

	Writers   []io.WriteCloser
	SkipClose bool

	Action  string
	Params  string
	Thought string

	Message *memory.Message `json:"message,omitempty"`
}

func NewActionRequest() *ActionRequest {
	return &ActionRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}

func (creq *ActionRequest) ToCompletionRequest() *CompletionRequest {
	req := NewCompletionRequest()
	req.Context = creq.Context
	req.Message = creq.Message
	req.Mode = "default"

	req.Writer = creq.Writers[0]

	return req
}

func (creq *ActionRequest) CountTokens() int {
	return creq.ToCompletionRequest().CountTokens()
}

func (creq *ActionRequest) ToChatCompletionMessages() []openai.ChatCompletionMessage {
	return creq.ToCompletionRequest().ToChatCompletionMessages()
}

func (creq *ActionRequest) GetID() string {
	return creq.ID
}

func (creq *ActionRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *ActionRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *ActionRequest) GetWriters() []io.WriteCloser {
	return creq.Writers
}

func (creq *ActionRequest) GetSkipClose() bool {
	return creq.SkipClose
}
