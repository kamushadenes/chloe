package cli

type ForgetCmd struct {
	All bool `help:"Forget all users, not just the CLI user"`
}

func (c *ForgetCmd) Run(globals *Globals) error {
	return Forget(globals.Context, c.All)
}
