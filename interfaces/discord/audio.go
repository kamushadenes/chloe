package discord

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func transcribe(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetAudios() {

		req := structs.NewActionRequest()
		req.Action = "transcribe"
		req.Params = path
		req.Message = msg
		req.Context = ctx
		req.Writer = NewDiscordWriter(ctx, req, true)

		if err := channels.RunAction(req); err != nil {
			return err
		}
	}

	return nil
}

func tts(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Action = "tts"
	req.Params = promptFromMessage(msg)
	req.Message = msg
	req.Context = ctx
	req.Writer = NewDiscordWriter(ctx, req, false)

	return channels.RunAction(req)
}
