package main

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/interfaces/cli"
	"github.com/kamushadenes/chloe/server"
	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func wait(quitCh chan os.Signal, errorCh chan error, cancel context.CancelFunc) {
	for {
		select {
		case <-quitCh:
			cancel()
			os.Exit(0)
		case err := <-errorCh:
			if err != nil {
				fmt.Println()
				fmt.Println(err)
			}
			cancel()
			os.Exit(1)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	errorCh := make(chan error)

	go wait(quitCh, errorCh, cancel)

	readyCh := make(chan bool)

	if len(os.Args) > 1 {
		flags.InteractiveCLI = isatty.IsTerminal(os.Stdout.Fd())

		zerolog.SetGlobalLevel(zerolog.Disabled)

		go server.InitServer(ctx, true, readyCh)

		<-readyCh

		errorCh <- cli.Handle(ctx)
	} else {
		fmt.Println("Chloe is starting...")
		go server.InitServer(ctx, false, readyCh)
		<-readyCh
	}
}
