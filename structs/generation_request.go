package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
)

type GenerationRequest struct {
	ID      string
	Context context.Context

	Writer    ChloeWriter
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message *memory.Message `json:"message,omitempty"`

	Prompt    string `json:"prompt"`
	Size      string `json:"size"`
	ImagePath string `json:"image"`
	Count     int    `json:"count"`
}

func NewGenerationRequest() *GenerationRequest {
	return &GenerationRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}
