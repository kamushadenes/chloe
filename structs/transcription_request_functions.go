package structs

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
)

func (creq *TranscriptionRequest) GetID() string {
	return creq.ID
}

func (creq *TranscriptionRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *TranscriptionRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *TranscriptionRequest) GetWriter() ChloeWriter {
	return creq.Writer
}

func (creq *TranscriptionRequest) SetWriter(w ChloeWriter) {
	creq.Writer = w
}

func (creq *TranscriptionRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *TranscriptionRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *TranscriptionRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *TranscriptionRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *TranscriptionRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
