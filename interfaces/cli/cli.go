package cli

import (
	"context"
	"github.com/alecthomas/kong"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/memory"
)

var user *memory.User

func Handle(ctx context.Context) error {
	var err error

	user, err = memory.GetUserByExternalID(ctx, "cli", "cli")
	if err != nil {
		user, err = memory.CreateUser(ctx, "User", "CLI", "cli")
		if err != nil {
			return err
		}
		err = user.AddExternalID(ctx, "cli", "cli")
		if err != nil {
			return err
		}
	}

	kongCtx := kong.Parse(&CLIFlags,
		kong.Name("chloe"),
		kong.Description("Chloe is a powerful AI Assistant\n\nRunning Chloe without arguments will start the server"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": flags.Version,
		})

	return kongCtx.Run(&Globals{Context: ctx})
}

func Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
