package discord

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func aiTranscribe(ctx context.Context, msg *memory.Message) error {
	for _, path := range msg.GetAudios() {

		req := structs.NewActionRequest()
		req.Action = "transcribe"
		req.Params = path
		req.Message = msg
		req.Context = ctx
		req.Writers = append(req.Writers, NewTextWriter(ctx, req, true))

		channels.ActionRequestsCh <- req
	}

	return nil
}

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
