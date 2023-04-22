package openai

import (
	"github.com/kamushadenes/chloe/langchain/llm/common"
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/sashabaranov/go-openai"
)

var (
	Ada = &common.LLMModel{
		Name:             openai.GPT3Ada,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0004, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	TextAda001 = &common.LLMModel{
		Name:             openai.GPT3TextAda001,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0004, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	Babbage = &common.LLMModel{
		Name:             openai.GPT3Babbage,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0005, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	TextBabbage001 = &common.LLMModel{
		Name:             openai.GPT3TextBabbage001,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0005, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	Curie = &common.LLMModel{
		Name:             openai.GPT3Curie,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0020, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	TextCurie001 = &common.LLMModel{
		Name:             openai.GPT3TextCurie001,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0020, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	Davinci = &common.LLMModel{
		Name:             openai.GPT3Davinci,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0200, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	TextDavinci002 = &common.LLMModel{
		Name:             openai.GPT3TextDavinci002,
		ContextSize:      4097,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0200, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	TextDavinci003 = &common.LLMModel{
		Name:             openai.GPT3TextDavinci003,
		ContextSize:      4097,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0200, Unit: schema.CostUnitToken, UnitSize: 1000},
	}
)
