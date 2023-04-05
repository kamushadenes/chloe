package utils

import (
	"time"
)

func TickerOrDefault(d time.Duration, def time.Duration) *time.Ticker {
	if d > 0 {
		return time.NewTicker(d)
	}

	return time.NewTicker(def)
}
