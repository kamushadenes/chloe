package structs

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type CalculationRequest struct {
	Context context.Context

	Writer    io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	User    *memory.User `json:"user,omitempty"`
	Content string       `json:"content"`
}

func (creq *CalculationRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *CalculationRequest) GetWriters() []io.WriteCloser {
	return []io.WriteCloser{creq.Writer}
}

func (creq *CalculationRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *CalculationRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *CalculationRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *CalculationRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *CalculationRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
