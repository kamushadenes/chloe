package structs

import (
	"context"
	"github.com/kamushadenes/chloe/users"
	"io"
)

type TTSRequest struct {
	Context context.Context

	Writers   []io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	User    *users.User `json:"user,omitempty"`
	Content string      `json:"content"`
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
