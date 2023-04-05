package cli

import (
	"fmt"
	"github.com/kamushadenes/chloe/memory"
)

type CreateAPIKeyCmd struct {
	UserID uint `arg:"" short:"u" long:"user-id" description:"User ID"`
}

func (c *CreateAPIKeyCmd) Run(globals *Globals) error {
	u, err := memory.GetUser(globals.Context, c.UserID)
	if err != nil {
		return err
	}

	key, err := u.CreateAPIKey(globals.Context)
	if err != nil {
		return err
	}

	fmt.Println(key)

	return nil
}
