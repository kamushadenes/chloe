package discord

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/actions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

func generate(ctx context.Context, msg *memory.Message) error {
	req := action_structs.NewActionRequest()
	req.Message = msg
	req.Context = ctx
	req.Action = "generate"
	req.Params["prompt"] = promptFromMessage(msg)
	req.Writer = NewDiscordWriter(ctx, req, false, req.Params["prompt"])
	req.Count = config.Discord.ImageCount

	return actions.HandleAction(req)
}
