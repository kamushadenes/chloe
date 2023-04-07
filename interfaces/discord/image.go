package discord

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func generate(ctx context.Context, msg *memory.Message) error {
	req := structs.NewActionRequest()
	req.Action = "generate"
	req.Params["prompt"] = promptFromMessage(msg)
	req.Message = msg
	req.Context = ctx
	req.Writer = NewDiscordWriter(ctx, req, false, req.Params["prompt"])
	req.Count = config.Discord.ImageCount

	return channels.RunAction(req)
}
