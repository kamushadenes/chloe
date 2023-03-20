package structs

import (
	"context"
	"github.com/kamushadenes/chloe/users"
	"io"
)

type TranscriptionRequest struct {
	Context context.Context

	Writer    io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	User     *users.User `json:"user,omitempty"`
	FilePath string      `json:"filePath"`
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
