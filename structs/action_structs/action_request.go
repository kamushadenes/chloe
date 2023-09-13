package action_structs

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/writer_structs"
)

type ActionRequest struct {
	ID      string
	Context context.Context

	Writer    writer_structs.ChloeWriter
	SkipClose bool

	ErrorChannel chan error

	Action string
	Params map[string]string

	Count int

	Message *memory.Message `json:"message,omitempty"`
}

func NewActionRequest() *ActionRequest {
	return &ActionRequest{
		ID:     uuid.Must(uuid.NewV4()).String(),
		Params: make(map[string]string),
	}
}
