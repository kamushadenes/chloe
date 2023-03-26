package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type ScrapeRequest struct {
	ID      string
	Context context.Context

	Writer    io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message *memory.Message `json:"message,omitempty"`

	Content string `json:"content"`
}

func NewScrapeRequest() *ScrapeRequest {
	return &ScrapeRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}

func (creq *ScrapeRequest) GetID() string {
	return creq.ID
}

func (creq *ScrapeRequest) GetMessage() *memory.Message {
	return creq.Message
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
