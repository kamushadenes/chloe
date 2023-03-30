package cli

import "github.com/kamushadenes/chloe/memory"

type MergeUsersCmd struct {
	Users []uint `arg:"" help:"Users to merge"`
}

func (c *MergeUsersCmd) Run(globals *Globals) error {
	return memory.MergeUsersByID(globals.Context, c.Users...)
}
