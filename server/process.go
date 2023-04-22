package server

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
)

func ProcessMessage(ctx context.Context, msg *memory.Message) error {
	logger := logging.FromContext(ctx)

	logger.Info().
		Uint("userID", msg.User.ID).
		Str("username", msg.User.Username).
		Str("firstName", msg.User.FirstName).
		Str("lastName", msg.User.LastName).
		Str("interface", msg.Interface).
		Str("content", msg.Content).
		Msg("message received")

	err := msg.Save(ctx)
	msg.ErrorCh <- err

	return err
}
