package cli

import (
	"context"
	"github.com/alecthomas/kong"
	"github.com/kamushadenes/chloe/flags"
)

func Handle(ctx context.Context) error {
	var err error

	user, err = getUser(ctx)
	if err != nil {
		return err
	}

	kongCtx := kong.Parse(&Flags,
		kong.Name("chloe"),
		kong.Description("Chloe is a powerful AI Assistant\n\nRunning Chloe without arguments will start the server"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
		kong.Vars{
			"version": flags.Version,
		})

	kongCtx.FatalIfErrorf(kongCtx.Run(&Globals{Context: ctx}))

	return nil
}

func Start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		}
	}
}
