package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
)

type TranscriptionRequest struct {
	ID      string
	Context context.Context

	Writer    ChloeWriter
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message  *memory.Message `json:"message,omitempty"`
	FilePath string          `json:"filePath"`
}

func NewTranscriptionRequest() *TranscriptionRequest {
	return &TranscriptionRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}
