package common

import (
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/utils"
)

type EmbeddingObject struct {
	Object    string    `json:"object"`
	Embedding []float32 `json:"embedding"`
	Index     int       `json:"index"`
}

type EmbeddingUsage struct {
	PromptTokens int
	TotalTokens  int
}

type EmbeddingCost struct {
	PromptCost float64
	TotalCost  float64
}

type EmbeddingResult struct {
	Embeddings []EmbeddingObject
	Cost       EmbeddingCost
	Usage      EmbeddingUsage
}

func (c *EmbeddingResult) CalculateCosts(m *EmbeddingModel) {
	c.Cost = EmbeddingCost{}

	c.Cost.PromptCost = m.UsageCost.Price * float64(c.Usage.PromptTokens) / float64(m.UsageCost.UnitSize)

	c.Cost.PromptCost = utils.RoundFloat(c.Cost.PromptCost, 6)
	c.Cost.TotalCost = c.Cost.PromptCost

	cost.AddCategoryCost("embedding", c.Cost.TotalCost)
}
