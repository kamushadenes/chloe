package utils

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/rs/zerolog"
	"time"
)

func WaitTimeout(ctx context.Context, ttype config.TimeoutType, fn func(ch chan interface{}, errCh chan error)) (interface{}, error) {
	logger := zerolog.Ctx(ctx)

	nch := make(chan interface{})
	errCh := make(chan error)

	go fn(nch, errCh)

	ticker := time.NewTicker(config.OpenAI.Timeouts[config.TimeoutTypeSlowness])
	timeout := time.NewTimer(config.OpenAI.Timeouts[ttype])

	for {
		select {
		case <-ticker.C:
			logger.Warn().Msg("still waiting for chain of thought analysis")
		case <-timeout.C:
			return nil, fmt.Errorf("timeout waiting for chain of thought analysis")
		case r := <-nch:
			ticker.Stop()
			timeout.Stop()

			return r, nil
		case err := <-errCh:
			ticker.Stop()
			timeout.Stop()

			return nil, err
		}
	}
}
