package cli

import "github.com/kamushadenes/chloe/memory"

type DeleteExternalIDCmd struct {
	UserID     uint   `short:"u" long:"user-id" description:"User ID"`
	ExternalID string `short:"e" long:"external-id" description:"External ID"`
	Interface  string `short:"i" long:"interface" description:"Interface"`
}

func (c *DeleteExternalIDCmd) Run(globals *Globals) error {
	u, err := memory.GetUser(globals.Context, c.UserID)
	if err != nil {
		return err
	}

	return u.DeleteExternalID(globals.Context, c.ExternalID, c.Interface)
}
