package common

import (
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/kamushadenes/chloe/utils"
	"time"
)

type ASRUsage struct {
	Duration time.Duration
}

type ASRCost struct {
	TotalCost float64
}

type ASRResult struct {
	Text  string `json:"text"`
	Usage ASRUsage
	Cost  ASRCost
}

func (c *ASRResult) CalculateCosts(m *ASRModel) {
	c.Cost = ASRCost{}

	if m.UsageCost != nil {
		switch m.UsageCost.Unit {
		case schema.CostUnitMinute:
			c.Cost.TotalCost = utils.RoundFloat(m.UsageCost.Price*float64(c.Usage.Duration.Minutes())/float64(m.UsageCost.UnitSize), 6)
		}
	}

	cost.AddCategoryCost("completion", c.Cost.TotalCost)
}
