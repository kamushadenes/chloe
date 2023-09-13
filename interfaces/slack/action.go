package slack

import (
	"context"
	"strings"

	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

func action(ctx context.Context, msg *memory.Message) error {
	fields := strings.Fields(msg.Content)

	req := action_structs.NewActionRequest()
	req.Context = ctx
	req.Action = fields[0]
	req.Params["text"] = strings.Join(fields[1:], " ")
	req.Writer = NewSlackWriter(ctx, req, false)

	//return structs.RunAction(req)
	return nil
}
