package common

import (
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/utils"
)

type LLMGeneration struct {
	Text         string
	Index        int
	FinishReason string
}

type LLMUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

type LLMCost struct {
	PromptCost     float64
	CompletionCost float64
	TotalCost      float64
}

type LLMResult struct {
	Generations []LLMGeneration
	Usage       LLMUsage
	Cost        LLMCost
}

func (c *LLMResult) CalculateCosts(m *LLMModel) {
	c.Cost = LLMCost{}

	if m.UsageCost != nil {
		c.Cost.PromptCost = m.UsageCost.Price * float64(c.Usage.PromptTokens) / float64(m.UsageCost.UnitSize)
		c.Cost.CompletionCost = m.UsageCost.Price * float64(c.Usage.CompletionTokens) / float64(m.UsageCost.UnitSize)
	} else if m.CompletionCost != nil && m.PromptCost != nil {
		c.Cost.PromptCost = m.PromptCost.Price * float64(c.Usage.PromptTokens) / float64(m.PromptCost.UnitSize)
		c.Cost.CompletionCost = m.CompletionCost.Price * float64(c.Usage.CompletionTokens) / float64(m.CompletionCost.UnitSize)
	}

	c.Cost.PromptCost = utils.RoundFloat(c.Cost.PromptCost, 6)
	c.Cost.CompletionCost = utils.RoundFloat(c.Cost.CompletionCost, 6)
	c.Cost.TotalCost = utils.RoundFloat(c.Cost.PromptCost+c.Cost.CompletionCost, 6)

	cost.AddCategoryCost("completion", c.Cost.TotalCost)
}
