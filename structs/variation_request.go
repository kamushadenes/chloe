package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
)

type VariationRequest struct {
	ID      string
	Context context.Context

	Writer    ChloeWriter
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message *memory.Message `json:"message,omitempty"`

	Size      string `json:"size"`
	ImagePath string `json:"image"`
	Count     int    `json:"count"`
}

func NewVariationRequest() *VariationRequest {
	return &VariationRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}
