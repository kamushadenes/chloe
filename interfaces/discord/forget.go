package discord

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func forgetUser(ctx context.Context, msg *memory.Message) error {
	return errors.Wrap(errors.ErrForgetUser, msg.User.DeleteMessages(ctx))
}
