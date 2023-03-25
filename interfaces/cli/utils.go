package cli

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/i18n"
)

func Forget(ctx context.Context) error {
	err := user.DeleteMessages(ctx)
	if err != nil {
		return err
	}

	fmt.Println(i18n.GetForgetText())

	return nil
}
