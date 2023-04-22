package structs

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func (creq *VariationRequest) GetID() string {
	return creq.ID
}

func (creq *VariationRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *VariationRequest) GetSize() string {
	return creq.Size
}

func (creq *VariationRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *VariationRequest) GetWriter() ChloeWriter {
	return creq.Writer
}

func (creq *VariationRequest) SetWriter(w ChloeWriter) {
	creq.Writer = w
}

func (creq *VariationRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *VariationRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *VariationRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *VariationRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *VariationRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
