package common

import (
	"context"
	"io"
)

type LLM interface {
	Generate(...string) (LLMResult, error)
	GenerateWithContext(context.Context, ...string) (LLMResult, error)
	GenerateWithOptions(context.Context, LLMOptions) (LLMResult, error)

	GenerateStream(io.Writer, ...string) (LLMResult, error)
	GenerateStreamWithContext(context.Context, io.Writer, ...string) (LLMResult, error)
	GenerateStreamWithOptions(context.Context, io.Writer, LLMOptions) (LLMResult, error)
}
