package openai

import (
	"github.com/kamushadenes/chloe/langchain/embeddings/common"
	"github.com/kamushadenes/chloe/langchain/schema"
)

var (
	AdaEmbeddingV2 = &common.EmbeddingModel{
		Name:        "text-embedding-ada-002",
		ContextSize: 8191,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.0001, Unit: schema.CostUnitToken, UnitSize: 1000},
	}
)
