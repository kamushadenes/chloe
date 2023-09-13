package slack

import (
	"context"

	"github.com/kamushadenes/chloe/langchain/actions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

func tts(ctx context.Context, msg *memory.Message) error {
	req := action_structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = "tts"
	req.Params["text"] = promptFromMessage(msg)
	req.Writer = NewSlackWriter(ctx, req, false, promptFromMessage(msg))

	return actions.HandleAction(req)
}
