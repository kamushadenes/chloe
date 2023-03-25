package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type CalculationRequest struct {
	ID      string
	Context context.Context

	Writer    io.WriteCloser
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message *memory.Message `json:"message,omitempty"`

	User    *memory.User `json:"user,omitempty"`
	Content string       `json:"content"`
}

func NewCalculationRequest() *CalculationRequest {
	return &CalculationRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}

func (creq *CalculationRequest) GetID() string {
	return creq.ID
}

func (creq *CalculationRequest) GetMessage() *memory.Message {
	return creq.Message
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
