package common

import (
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/utils"
)

type ChatGeneration struct {
	Text         string
	Message      messages.Message
	FinishReason string
}

type ChatUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

type ChatCost struct {
	PromptCost     float64
	CompletionCost float64
	TotalCost      float64
}

type ChatResult struct {
	Generations []ChatGeneration
	Usage       ChatUsage
	Cost        ChatCost
}

func (c *ChatResult) CalculateCosts(m *ChatModel) {
	c.Cost = ChatCost{}

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
