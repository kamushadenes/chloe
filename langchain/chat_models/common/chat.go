package common

import (
	"context"
	"io"
)

type Chat interface {
	Chat(...Message) (ChatResult, error)
	ChatWithContext(context.Context, ...Message) (ChatResult, error)
	ChatWithOptions(context.Context, ChatOptions) (ChatResult, error)

	ChatStream(io.Writer, ...Message) (ChatResult, error)
	ChatStreamWithContext(context.Context, io.Writer, ...Message) (ChatResult, error)
	ChatStreamWithOptions(context.Context, io.Writer, ChatOptions) (ChatResult, error)
}
