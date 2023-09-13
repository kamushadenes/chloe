package discord

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func forgetUser(ctx context.Context, msg *memory.Message) error {
	err := msg.User.DeleteMessages(ctx)
	if err != nil {
		return errors.Wrap(errors.ErrForgetUser, err)
	}

	return nil
}
