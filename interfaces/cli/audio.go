package cli

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
)

func TTS(ctx context.Context, text string, writer structs.ChloeWriter) error {
	req := structs.NewActionRequest()
	req.Context = ctx
	req.Message = memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "cli")
	req.Message.User = user

	req.Writer = writer

	req.Action = "tts"
	req.Params["text"] = text

	return channels.RunAction(req)
}
