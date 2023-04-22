package discord

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

func action(ctx context.Context, msg *memory.Message) error {
	fields := strings.Fields(msg.Content)

	req := structs.NewActionRequest()
	req.Context = ctx
	req.Message = msg
	req.Action = fields[0]
	req.Params["text"] = strings.Join(fields[1:], " ")
	req.Thought = fmt.Sprintf("User wants to run action %s", fields[0])
	req.Writer = NewDiscordWriter(ctx, req, false)
	req.Count = config.Discord.ImageCount

	return channels.RunAction(req)
}
