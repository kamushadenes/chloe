package discord

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func generate(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Action = "image"
	req.Params = promptFromMessage(msg)
	req.Message = msg
	req.Context = ctx
	req.Writers = append(req.Writers, NewImageWriter(ctx, req, false, req.Params))

	return channels.RunAction(req)
}
