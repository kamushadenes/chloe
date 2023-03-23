package structs

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type ScrapeRequest struct {
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

func (creq *ScrapeRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *ScrapeRequest) GetWriters() []io.WriteCloser {
	return []io.WriteCloser{creq.Writer}
}

func (creq *ScrapeRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *ScrapeRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *ScrapeRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *ScrapeRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *ScrapeRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
