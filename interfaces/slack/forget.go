package slack

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func forgetUser(ctx context.Context, msg *memory.Message) error {
	return msg.User.DeleteMessages(ctx)
}
