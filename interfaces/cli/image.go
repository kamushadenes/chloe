package cli

import (
	"context"
	"github.com/kamushadenes/chloe/providers/openai"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"os"
)

func Generate(ctx context.Context, text string) error {
	return openai.Generate(ctx, &structs.GenerationRequest{
		Context: ctx,
		Writers: []io.WriteCloser{os.Stdout},
		User:    user,
		Prompt:  text,
	})
}
