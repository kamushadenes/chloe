package cli

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/i18n"
	"github.com/kamushadenes/chloe/memory"
)

func Forget(ctx context.Context, all bool) error {
	var err error

	if all {
		err = memory.DeleteAllMessages(ctx)
	} else {
		err = user.DeleteMessages(ctx)
	}
	if err != nil {
		return err
	}

	fmt.Println(i18n.GetForgetText())

	return nil
}
