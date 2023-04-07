package cost

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"time"
)

func MonitorCost(ctx context.Context) {
	defer func() {
		logCost("total session cost")
	}()

	ticker := time.NewTicker(config.Misc.CostTrackingInterval)

	for {
		select {
		case <-ticker.C:
			logCost("session cost")
		case <-ctx.Done():
			return
		}
	}
}
