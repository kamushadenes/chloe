package common

import (
	"time"
)

type TTSUsage struct {
	Duration time.Duration
}

type TTSCost struct {
	TotalCost float64
}

type TTSResult struct {
	Audio       []byte
	ContentType string
	Usage       TTSUsage
	Cost        TTSCost
}

func (c *TTSResult) CalculateCosts() {
	c.Cost = TTSCost{}

	// cost.AddCategoryCost("completion", c.Cost.TotalCost)
}
