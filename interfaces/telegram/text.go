package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/messages"
	"github.com/kamushadenes/chloe/structs"
)

func aiComplete(ctx context.Context, msg *messages.Message, ch chan interface{}) error {
	request := &structs.CompletionRequest{}

	request.User = msg.User

	if request.Mode == "" {
		mode, _ := memory.GetUserMode(ctx, request.User.ID)
		request.Mode = mode
	}
	request.Args = map[string]interface{}{
		"User": msg.User,
		"Mode": request.Mode,
	}

	request.Content = msg.Source.Telegram.Update.Message.Text

	request.ResultChannel = ch
	request.Context = ctx
	request.Writer = NewTextWriter(ctx, msg, false)

	channels.CompletionRequestsCh <- request

	return nil
}
