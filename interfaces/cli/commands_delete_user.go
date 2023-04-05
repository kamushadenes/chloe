package cli

import "github.com/kamushadenes/chloe/memory"

type DeleteUserCmd struct {
	UserID uint `short:"u" long:"user-id" description:"User ID"`
}

func (c *DeleteUserCmd) Run(globals *Globals) error {
	u, err := memory.GetUser(globals.Context, c.UserID)
	if err != nil {
		return err
	}

	if err := u.DeleteMessages(globals.Context); err != nil {
		return err
	}

	return u.Delete(globals.Context)
}
