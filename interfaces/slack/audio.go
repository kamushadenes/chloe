package slack

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
)

func tts(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Action = "tts"
	req.Params["text"] = promptFromMessage(msg)
	req.Message = msg
	req.Context = ctx
	req.Writer = NewSlackWriter(ctx, req, false)

	return channels.RunAction(req)
}
