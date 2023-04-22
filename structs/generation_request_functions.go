package structs

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func (creq *GenerationRequest) GetID() string {
	return creq.ID
}

func (creq *GenerationRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *GenerationRequest) GetSize() string {
	return creq.Size
}

func (creq *GenerationRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *GenerationRequest) GetWriter() ChloeWriter {
	return creq.Writer
}

func (creq *GenerationRequest) SetWriter(w ChloeWriter) {
	creq.Writer = w
}

func (creq *GenerationRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *GenerationRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *GenerationRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *GenerationRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *GenerationRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
