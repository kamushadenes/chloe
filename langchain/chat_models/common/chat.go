package common

import (
	"context"
	"io"

	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
)

type Chat interface {
	Chat(...messages.Message) (ChatResult, error)
	ChatWithContext(context.Context, ...messages.Message) (ChatResult, error)
	ChatWithOptions(context.Context, ChatOptions) (ChatResult, error)

	ChatStream(io.Writer, ...messages.Message) (ChatResult, error)
	ChatStreamWithContext(context.Context, io.Writer, ...messages.Message) (ChatResult, error)
	ChatStreamWithOptions(context.Context, io.Writer, ChatOptions) (ChatResult, error)
}
