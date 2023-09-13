package action_structs

import (
	"context"

	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/completion_request_structs"
	"github.com/kamushadenes/chloe/structs/writer_structs"
	"github.com/sashabaranov/go-openai"
)

func (creq *ActionRequest) ToCompletionRequest() *completion_request_structs.CompletionRequest {
	req := completion_request_structs.NewCompletionRequest()
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

func (creq *ActionRequest) GetWriter() writer_structs.ChloeWriter {
	return creq.Writer
}

func (creq *ActionRequest) SetWriter(w writer_structs.ChloeWriter) {
	creq.Writer = w
}

func (creq *ActionRequest) GetSkipClose() bool {
	return creq.SkipClose
}
