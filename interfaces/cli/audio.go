package cli

import (
	"context"
	"github.com/kamushadenes/chloe/providers/google"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"os"
)

func TTS(ctx context.Context, text string) error {
	return google.TTS(ctx, &structs.TTSRequest{
		Context: ctx,
		User:    user,
		Writers: []io.WriteCloser{os.Stdout},
		Content: text,
	})
}
