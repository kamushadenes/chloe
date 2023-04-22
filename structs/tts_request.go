package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/langchain/memory"
)

type TTSRequest struct {
	ID      string
	Context context.Context

	Writer    ChloeWriter
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message *memory.Message `json:"message,omitempty"`

	Content string `json:"content"`
}

func NewTTSRequest() *TTSRequest {
	return &TTSRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}
