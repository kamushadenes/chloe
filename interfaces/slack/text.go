package slack

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func aiComplete(ctx context.Context, msg *memory.Message) error {
	request := structs.NewCompletionRequest()

	if request.Mode == "" {
		request.Mode = msg.User.Mode
	}
	request.Args = map[string]interface{}{
		"User": msg.User,
		"Mode": request.Mode,
	}

	request.Message = msg

	request.Context = ctx
	request.Writer = NewTextWriter(ctx, request, false)

	channels.CompletionRequestsCh <- request

	return nil
}
