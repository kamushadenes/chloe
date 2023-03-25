package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type TranscriptionRequest struct {
	ID      string
	Context context.Context

	Writer    io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	User     *memory.User    `json:"user,omitempty"`
	Message  *memory.Message `json:"message,omitempty"`
	FilePath string          `json:"filePath"`
}

func NewTranscriptionRequest() *TranscriptionRequest {
	return &TranscriptionRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}

func (creq *TranscriptionRequest) GetID() string {
	return creq.ID
}

func (creq *TranscriptionRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *TranscriptionRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *TranscriptionRequest) GetWriters() []io.WriteCloser {
	return []io.WriteCloser{creq.Writer}
}

func (creq *TranscriptionRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *TranscriptionRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *TranscriptionRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *TranscriptionRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *TranscriptionRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
