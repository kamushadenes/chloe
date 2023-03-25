package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type TTSRequest struct {
	ID      string
	Context context.Context

	Writers   []io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message *memory.Message `json:"message,omitempty"`

	User    *memory.User `json:"user,omitempty"`
	Content string       `json:"content"`
}

func NewTTSRequest() *TTSRequest {
	return &TTSRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}

func (creq *TTSRequest) GetID() string {
	return creq.ID
}

func (creq *TTSRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *TTSRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *TTSRequest) GetWriters() []io.WriteCloser {
	return creq.Writers
}

func (creq *TTSRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *TTSRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *TTSRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *TTSRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *TTSRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
