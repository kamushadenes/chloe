package cli

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
)

var user *memory.User

func getUser(ctx context.Context) (*memory.User, error) {
	u, err := memory.GetUserByExternalID(ctx, "cli", "cli")
	if err != nil {
		u, err = memory.CreateUser(ctx, "User", "CLI", "cli")
		if err != nil {
			return nil, err
		}
		err = u.AddExternalID(ctx, "cli", "cli")
		if err != nil {
			return nil, err
		}
	}

	return u, err
}
