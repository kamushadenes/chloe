package cli

import (
	"context"
	"fmt"
)

func Forget(ctx context.Context) error {
	err := user.DeleteMessages(ctx)
	if err != nil {
		return err
	}

	fmt.Println("Forgot")

	return nil
}
