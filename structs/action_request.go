package structs

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/memory"
)

type ActionRequest struct {
	ID      string
	Context context.Context

	Writer    ChloeWriter
	SkipClose bool

	ErrorChannel chan error

	Action  string
	Params  string
	Thought string

	Message *memory.Message `json:"message,omitempty"`
}

func NewActionRequest() *ActionRequest {
	return &ActionRequest{
		ID: uuid.Must(uuid.NewV4()).String(),
	}
}
