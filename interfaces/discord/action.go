package discord

import (
	"context"
	"strings"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

func action(ctx context.Context, msg *memory.Message) error {
	fields := strings.Fields(msg.Content)

	req := action_structs.NewActionRequest()
	req.Context = ctx
	req.Message = msg
	req.Action = fields[0]
	req.Params["text"] = strings.Join(fields[1:], " ")
	req.Writer = NewDiscordWriter(ctx, req, false)
	req.Count = config.Discord.ImageCount

	//return structs.RunAction(req)
	return nil
}
