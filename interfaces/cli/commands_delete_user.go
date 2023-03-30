package cli

import "github.com/kamushadenes/chloe/memory"

type DeleteUserCmd struct {
	UserID uint `short:"u" long:"user-id" description:"User ID"`
}

func (c *DeleteUserCmd) Run(globals *Globals) error {
	user, err := memory.GetUser(globals.Context, c.UserID)
	if err != nil {
		return err
	}

	if err := user.DeleteMessages(globals.Context); err != nil {
		return err
	}

	return user.Delete(globals.Context)
}
