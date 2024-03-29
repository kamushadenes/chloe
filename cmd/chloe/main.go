package main

import (
	"context"
	"fmt"
	"os"

	"github.com/pkg/profile"

	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/interfaces/cli"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/server"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
)

var prf interface{ Stop() }

func wait(quitCh chan os.Signal, errorCh chan error, cancel context.CancelFunc) {
	for {
		select {
		case <-quitCh:
			cancel()
			prf.Stop()
			os.Exit(0)
		case err := <-errorCh:
			if err != nil {
				fmt.Println()
				fmt.Println(err)
			}
			cancel()
			prf.Stop()
			os.Exit(1)
		}
	}
}

func main() {
	prf = profile.Start(profile.CPUProfile, profile.ProfilePath("."))
	defer prf.Stop()

	flags.Debug = os.Getenv("DEBUG") == "true"

	ctx, cancel := context.WithCancel(context.Background())

	quitCh := make(chan os.Signal, 1)
	/*signal.Notify(quitCh,
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGINT,
		syscall.SIGQUIT,
	)*/

	errorCh := make(chan error)

	go wait(quitCh, errorCh, cancel)

	readyCh := make(chan bool)

	if len(os.Args) > 1 {
		flags.InteractiveCLI = isatty.IsTerminal(os.Stdout.Fd())

		if flags.InteractiveCLI {
			logging.Disable()
		}

		go server.InitServer(ctx, true, readyCh)

		<-readyCh

		zerolog.SetGlobalLevel(zerolog.Disabled)

		errorCh <- cli.Handle(ctx)
	} else {
		fmt.Println("Chloe is starting...")
		go server.InitServer(ctx, false, readyCh)
		<-readyCh
	}
}
