package openai

import (
	"github.com/kamushadenes/chloe/langchain/embeddings/common"
	"github.com/kamushadenes/chloe/langchain/schema"
)

var (
	AdaSimilarity = &common.EmbeddingModel{
		Name:        "text-similarity-ada-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.004, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	BabbageSimilarity = &common.EmbeddingModel{
		Name:        "text-similarity-babbage-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.005, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	CurieSimilarity = &common.EmbeddingModel{
		Name:        "text-similarity-curie-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.02, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	DavinciSimilarity = &common.EmbeddingModel{
		Name:        "text-similarity-davinci-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.2, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	AdaSearchDocument = &common.EmbeddingModel{
		Name:        "text-search-ada-doc-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.004, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	BabbageSearchDocument = &common.EmbeddingModel{
		Name:        "text-search-babbage-doc-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.005, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	CurieSearchDocument = &common.EmbeddingModel{
		Name:        "text-search-curie-doc-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.02, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	DavinciSearchDocument = &common.EmbeddingModel{
		Name:        "text-search-davinci-doc-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.2, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	AdaSearchQuery = &common.EmbeddingModel{
		Name:        "text-search-ada-query-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.004, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	BabbageSearchQuery = &common.EmbeddingModel{
		Name:        "text-search-babbage-query-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.005, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	CurieSearchQuery = &common.EmbeddingModel{
		Name:        "text-search-curie-query-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.02, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	DavinciSearchQuery = &common.EmbeddingModel{
		Name:        "text-search-davinci-query-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.2, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	AdaCodeSearchCode = &common.EmbeddingModel{
		Name:        "code-search-ada-code-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.004, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	BabbageCodeSearchCode = &common.EmbeddingModel{
		Name:        "code-search-babbage-code-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.005, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	AdaCodeSearchText = &common.EmbeddingModel{
		Name:        "code-search-ada-text-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.004, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	BabbageCodeSearchText = &common.EmbeddingModel{
		Name:        "code-search-babbage-text-001",
		ContextSize: 2046,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.005, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	AdaEmbeddingV2 = &common.EmbeddingModel{
		Name:        "text-embedding-ada-002",
		ContextSize: 8191,
		ContextUnit: schema.ContextUnitToken,
		UsageCost:   &schema.CostObject{Price: 0.0001, Unit: schema.CostUnitToken, UnitSize: 1000},
	}
)
