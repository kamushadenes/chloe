package completion_request_structs

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/writer_structs"
)

type CompletionRequest struct {
	ID      string
	Context context.Context

	Writer    writer_structs.ChloeWriter
	SkipClose bool

	StartChannel    chan bool
	ContinueChannel chan bool
	ErrorChannel    chan error
	ResultChannel   chan interface{}

	Message *memory.Message

	Mode string                 `json:"mode"`
	Args map[string]interface{} `json:"args"`
}

func NewCompletionRequest() *CompletionRequest {
	return &CompletionRequest{
		ID:   uuid.Must(uuid.NewV4()).String(),
		Args: make(map[string]interface{}),
	}
}
