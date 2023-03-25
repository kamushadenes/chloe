package utils

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/rs/zerolog"
	"time"
)

func WaitTimeout(ctx context.Context, timeout time.Duration, fn func(ch chan interface{}, errCh chan error)) (interface{}, error) {
	logger := zerolog.Ctx(ctx)

	nch := make(chan interface{})
	errCh := make(chan error)

	go fn(nch, errCh)

	ticker := time.NewTicker(config.Timeouts.SlownessWarning)
	timeoutTicker := time.NewTimer(timeout)

	for {
		select {
		case <-ticker.C:
			logger.Warn().Msg("still waiting for request to complete")
		case <-timeoutTicker.C:
			logger.Debug().Msg("request timed out")
			return nil, fmt.Errorf("timeout waiting for request to complete")
		case r := <-nch:
			logger.Debug().Msg("finished waiting for request to complete")
			ticker.Stop()
			timeoutTicker.Stop()

			return r, nil
		case err := <-errCh:
			logger.Debug().Msg("request completed with error")
			ticker.Stop()
			timeoutTicker.Stop()

			return nil, err
		}
	}
}
