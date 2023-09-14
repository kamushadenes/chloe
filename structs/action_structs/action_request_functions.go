package action_structs

import (
	"context"

	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/writer_structs"
)

func (creq *ActionRequest) GetID() string {
	return creq.ID
}

func (creq *ActionRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *ActionRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *ActionRequest) GetWriter() writer_structs.ChloeWriter {
	return creq.Writer
}

func (creq *ActionRequest) SetWriter(w writer_structs.ChloeWriter) {
	creq.Writer = w
}

func (creq *ActionRequest) GetSkipClose() bool {
	return creq.SkipClose
}
