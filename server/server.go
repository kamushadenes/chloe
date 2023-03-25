package server

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/interfaces/cli"
	"github.com/kamushadenes/chloe/interfaces/http"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/providers/openai"
)

func InitServer(ctx context.Context, isCLI bool, readyCh chan bool) {
	ctx, cancel := context.WithCancel(ctx)

	logger := logging.GetLogger()

	if config.OpenAI.APIKey == "" {
		logger.Fatal().Msg("OpenAI API key is not set")
	}

	_, err := memory.Setup(ctx)
	if err != nil {
		panic(err)
	}

	go memory.Start(ctx)

	go MonitorMessages(ctx)
	go MonitorRequests(ctx)

	go openai.MonitorSummary(ctx)
	go openai.MonitorModeration(ctx)

	if isCLI {
		go cli.Start(ctx)
		readyCh <- true
	} else {
		go http.Start(ctx)
		go telegram.Start(ctx)
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
