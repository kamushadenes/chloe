package cli

import "github.com/kamushadenes/chloe/memory"

type AddExternalIDCmd struct {
	UserID     uint   `short:"u" long:"user-id" description:"User ID"`
	ExternalID string `short:"e" long:"external-id" description:"External ID"`
	Interface  string `short:"i" long:"interface" description:"Interface"`
}

func (c *AddExternalIDCmd) Run(globals *Globals) error {
	user, err := memory.GetUser(globals.Context, c.UserID)
	if err != nil {
		return err
	}

	return user.AddExternalID(globals.Context, c.ExternalID, c.Interface)
}
