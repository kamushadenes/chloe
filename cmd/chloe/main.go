package main

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/interfaces/cli"
	"github.com/kamushadenes/chloe/server"
	"github.com/rs/zerolog"
	"os"
	"os/signal"
	"syscall"
)

func wait(quitCh chan os.Signal, errorCh chan bool, cancel context.CancelFunc) {
	for {
		select {
		case <-quitCh:
			cancel()
			os.Exit(0)
		case <-errorCh:
			cancel()
			os.Exit(1)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	quitCh := make(chan os.Signal)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	errorCh := make(chan bool)

	go wait(quitCh, errorCh, cancel)

	readyCh := make(chan bool)

	if len(os.Args) > 1 {
		zerolog.SetGlobalLevel(zerolog.Disabled)

		go server.InitServer(ctx, true, readyCh)

		<-readyCh

		if err := cli.Handle(ctx); err != nil {
			errorCh <- true
		}
	} else {
		fmt.Println("Chloe is starting...")
		go server.InitServer(ctx, false, readyCh)
		<-readyCh
	}
}
