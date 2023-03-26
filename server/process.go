package server

import (
	"context"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
)

func ProcessMessage(ctx context.Context, msg *memory.Message) error {
	logger := logging.GetLogger()

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
