package cli

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/i18n"
)

func Forget(ctx context.Context, all bool) error {
	var err error

	if all {
		err = user.DeleteAllMessages(ctx)
	} else {
		err = user.DeleteMessages(ctx)
	}
	if err != nil {
		return err
	}

	fmt.Println(i18n.GetForgetText())

	return nil
}
