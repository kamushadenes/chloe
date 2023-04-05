package slack

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

func action(ctx context.Context, msg *memory.Message) error {
	fields := strings.Fields(msg.Content)

	req := structs.NewActionRequest()
	req.Context = ctx
	req.Action = fields[0]
	req.Params = strings.Join(fields[1:], " ")
	req.Thought = fmt.Sprintf("User wants to run action %s", fields[0])
	req.Writer = NewSlackWriter(ctx, req, false)

	return channels.RunAction(req)
}
