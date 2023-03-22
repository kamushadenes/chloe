package cli

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/memory"
)

func Forget(ctx context.Context) error {
	err := memory.DeleteMessages(ctx, user.ID)
	if err != nil {
		return err
	}

	fmt.Println("Forgot")

	return nil
}
