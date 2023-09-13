package openai

import (
	"github.com/kamushadenes/chloe/langchain/diffusion_models/common"
	"github.com/kamushadenes/chloe/langchain/schema"
)

var (
	DallE256X256 = &common.DiffusionModel{
		Name:        "dall-e-256x256",
		ContextSize: 1000,
		ContextUnit: schema.ContextUnitChar,
		UsageCost:   &schema.CostObject{Price: 0.016, Unit: schema.CostUnitImage, UnitSize: 1},
	}

	DallE512X512 = &common.DiffusionModel{
		Name:        "dall-e-512x512",
		ContextSize: 1000,
		ContextUnit: schema.ContextUnitChar,
		UsageCost:   &schema.CostObject{Price: 0.018, Unit: schema.CostUnitImage, UnitSize: 1},
	}

	DallE1024X1024 = &common.DiffusionModel{
		Name:        "dall-e-1024x1024",
		ContextSize: 1000,
		ContextUnit: schema.ContextUnitChar,
		UsageCost:   &schema.CostObject{Price: 0.020, Unit: schema.CostUnitImage, UnitSize: 1},
	}
)
