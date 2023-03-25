package cli

import (
	"context"
	"github.com/kamushadenes/chloe/providers/openai"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"os"
)

func Generate(ctx context.Context, text string) error {
	req := structs.NewGenerationRequest()
	req.Context = ctx
	req.User = user
	req.Writers = []io.WriteCloser{os.Stdout}
	req.Prompt = text

	return openai.Generate(ctx, req)
}
