package chat_models

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/schema"
	"io"
)

type Chat interface {
	Chat(...schema.Message) (schema.ChatResult, error)
	ChatWithContext(context.Context, ...schema.Message) (schema.ChatResult, error)
	ChatWithOptions(context.Context, schema.ChatOptions) (schema.ChatResult, error)

	ChatStream(io.Writer, ...schema.Message) (schema.ChatResult, error)
	ChatStreamWithContext(context.Context, io.Writer, ...schema.Message) (schema.ChatResult, error)
	ChatStreamWithOptions(context.Context, io.Writer, schema.ChatOptions) (schema.ChatResult, error)
}
