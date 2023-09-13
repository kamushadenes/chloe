package common

import (
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/utils"
)

type DiffusionUsage struct {
	ImageCount int
}

type DiffusionCost struct {
	TotalCost float64
}

type DiffusionResult struct {
	Images [][]byte
	Usage  DiffusionUsage
	Cost   DiffusionCost
}

func (c *DiffusionResult) CalculateCosts(m *DiffusionModel) {
	c.Cost = DiffusionCost{}

	if m.UsageCost != nil {
		c.Cost.TotalCost = utils.RoundFloat(m.UsageCost.Price*float64(c.Usage.ImageCount)/float64(m.UsageCost.UnitSize), 6)

		cost.AddCategoryCost("completion", c.Cost.TotalCost)
	}
}
