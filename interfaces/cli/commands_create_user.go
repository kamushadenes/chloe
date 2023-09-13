package cli

import (
	"fmt"
	"github.com/kamushadenes/chloe/langchain/memory"
)

type CreateUserCmd struct {
	Username  string `short:"u" long:"username" description:"Username"`
	FirstName string `short:"f" long:"first-name" description:"First name"`
	LastName  string `short:"l" long:"last-name" description:"Last name"`
}

func (c *CreateUserCmd) Run(globals *Globals) error {
	u, err := memory.CreateUser(globals.Context, c.Username, c.FirstName, c.LastName)
	if err != nil {
		return err
	}

	fmt.Printf("User created: %d\n", u.ID)

	return nil
}
