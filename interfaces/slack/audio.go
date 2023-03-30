package slack

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func aiTTS(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Action = "tts"
	req.Params = promptFromMessage(msg)
	req.Message = msg
	req.Context = ctx
	req.Writers = append(req.Writers, NewAudioWriter(ctx, req, false))

	channels.ActionRequestsCh <- req

	return nil
}
