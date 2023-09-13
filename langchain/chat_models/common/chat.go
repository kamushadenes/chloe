package common

import (
	"context"
	"io"

	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
)

type Chat interface {
	Chat(*memory.Message, ...messages.Message) (ChatResult, error)
	ChatWithContext(context.Context, *memory.Message, ...messages.Message) (ChatResult, error)
	ChatWithOptions(context.Context, ChatOptions) (ChatResult, error)

	ChatStream(io.Writer, *memory.Message, ...messages.Message) (ChatResult, error)
	ChatStreamWithContext(context.Context, io.Writer, *memory.Message, ...messages.Message) (ChatResult, error)
	ChatStreamWithOptions(context.Context, io.Writer, ChatOptions) (ChatResult, error)
}
