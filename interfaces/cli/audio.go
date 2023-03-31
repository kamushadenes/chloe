package cli

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"os"
)

func TTS(ctx context.Context, text string, writers ...io.WriteCloser) error {
	req := structs.NewActionRequest()
	req.Context = ctx
	req.Message = memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "cli")
	req.Message.User = user
	if len(writers) > 0 {
		req.Writers = writers
	} else {
		req.Writers = []io.WriteCloser{os.Stdout}
	}
	req.Action = "tts"
	req.Params = text

	return channels.RunAction(req)
}
