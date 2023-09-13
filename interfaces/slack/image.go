package slack

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/actions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

func generate(ctx context.Context, msg *memory.Message) error {
	req := action_structs.NewActionRequest()
	req.Context = ctx
	req.Message = msg
	req.Action = "generate"
	req.Params["prompt"] = promptFromMessage(msg)
	req.Writer = NewSlackWriter(ctx, req, false, promptFromMessage(msg))
	req.Count = config.Slack.ImageCount

	return actions.HandleAction(req)
}
