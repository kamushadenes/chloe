package cli

import (
	"context"
	"github.com/kamushadenes/chloe/providers/google"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"os"
)

func TTS(ctx context.Context, text string) error {
	req := structs.NewTTSRequest()
	req.Context = ctx
	req.User = user
	req.Writers = []io.WriteCloser{os.Stdout}
	req.Content = text

	return google.TTS(ctx, req)
}
