package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

func aiComplete(ctx context.Context, msg *memory.Message, ch chan interface{}) error {
	request := structs.NewCompletionRequest()

	if request.Mode == "" {
		request.Mode = msg.User.Mode
	}
	request.Args = map[string]interface{}{
		"User": msg.User,
		"Mode": request.Mode,
	}

	request.Message = msg

	request.ResultChannel = ch
	request.Context = ctx
	request.Writer = NewTelegramWriter(ctx, request, false)

	return channels.RunCompletion(request)
}
