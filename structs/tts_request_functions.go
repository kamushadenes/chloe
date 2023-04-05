package structs

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
)

func (creq *TTSRequest) GetID() string {
	return creq.ID
}

func (creq *TTSRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *TTSRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *TTSRequest) GetWriter() ChloeWriter {
	return creq.Writer
}

func (creq *TTSRequest) SetWriter(w ChloeWriter) {
	creq.Writer = w
}

func (creq *TTSRequest) GetSkipClose() bool {
	return creq.SkipClose
}

func (creq *TTSRequest) GetStartChannel() chan bool {
	return creq.StartChannel
}

func (creq *TTSRequest) GetContinueChannel() chan bool {
	return creq.ContinueChannel
}

func (creq *TTSRequest) GetErrorChannel() chan error {
	return creq.ErrorChannel
}

func (creq *TTSRequest) GetResultChannel() chan interface{} {
	return creq.ResultChannel
}
