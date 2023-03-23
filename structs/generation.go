package structs

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type GenerationRequest struct {
	Context context.Context

	Writers   []io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	User      *memory.User `json:"user,omitempty"`
	Prompt    string       `json:"prompt"`
	Size      string       `json:"size"`
	ImagePath string       `json:"image"`
}

func (creq *GenerationRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *GenerationRequest) GetWriters() []io.WriteCloser {
	return creq.Writers
}

func (creq *GenerationRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *GenerationRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *GenerationRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *GenerationRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *GenerationRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
