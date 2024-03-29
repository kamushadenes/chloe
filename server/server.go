package server

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/interfaces/cli"
	"github.com/kamushadenes/chloe/interfaces/discord"
	"github.com/kamushadenes/chloe/interfaces/http"
	"github.com/kamushadenes/chloe/interfaces/slack"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
)

func InitServer(ctx context.Context, isCLI bool, readyCh chan bool) {
	ctx, cancel := context.WithCancel(ctx)

	logger := logging.FromContext(ctx)

	if config.OpenAI.APIKey == "" {
		logger.Fatal().Msg("OpenAI API key is not set")
	}

	if _, err := memory.Setup(ctx); err != nil {
		panic(err)
	}

	go memory.Start(ctx)

	go cost.MonitorCost(ctx)

	if isCLI {
		go cli.Start(ctx)
		readyCh <- true
	} else {
		go http.Start(ctx)
		go telegram.Start(ctx)
		go discord.Start(ctx)
		go slack.Start(ctx)
	}

	for {
		select {
		case <-ctx.Done():
			readyCh <- true
			cancel()
			return
		}
	}
}
