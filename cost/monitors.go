package cost

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"time"
)

func logCost(msg string) {
	logger := logging.GetLogger()

	l := logger.Info()

	for _, k := range GetCategories() {
		l = l.Float64(k, GetCategoryCost(k))
	}

	l = l.Float64("total", GetTotalSessionCost())

	l.Msg(msg)
}

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
