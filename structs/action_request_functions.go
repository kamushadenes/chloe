package structs

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"github.com/sashabaranov/go-openai"
)

func (creq *ActionRequest) ToCompletionRequest() *CompletionRequest {
	req := NewCompletionRequest()
	req.Context = creq.Context
	req.Message = creq.Message
	req.Mode = "default"

	req.Writer = creq.Writer

	return req
}

func (creq *ActionRequest) CountTokens() int {
	return creq.ToCompletionRequest().CountTokens()
}

func (creq *ActionRequest) ToChatCompletionMessages() []openai.ChatCompletionMessage {
	return creq.ToCompletionRequest().ToChatCompletionMessages()
}

func (creq *ActionRequest) GetID() string {
	return creq.ID
}

func (creq *ActionRequest) GetMessage() *memory.Message {
	return creq.Message
}

func (creq *ActionRequest) GetContext() context.Context {
	return creq.Context
}

func (creq *ActionRequest) GetWriter() ChloeWriter {
	return creq.Writer
}

func (creq *ActionRequest) SetWriter(w ChloeWriter) {
	creq.Writer = w
}

func (creq *ActionRequest) GetSkipClose() bool {
	return creq.SkipClose
}
