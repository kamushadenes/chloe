package structs

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type VariationRequest struct {
	Context context.Context

	Writers   []io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	User      *memory.User `json:"user,omitempty"`
	Size      string       `json:"size"`
	ImagePath string       `json:"image"`
}

func (creq *VariationRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *VariationRequest) GetWriters() []io.WriteCloser {
	return creq.Writers
}

func (creq *VariationRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *VariationRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *VariationRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *VariationRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *VariationRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
