package openai

import (
	"github.com/kamushadenes/chloe/langchain/llm/common"
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/sashabaranov/go-openai"
)

var (
	Babbage = &common.LLMModel{
		Name:             openai.GPT3Babbage,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0005, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	Davinci = &common.LLMModel{
		Name:             openai.GPT3Davinci,
		ContextSize:      2049,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0200, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT3Dot5TurboInstruct = &common.LLMModel{
		Name:             openai.GPT3Dot5TurboInstruct,
		ContextSize:      4097,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.0012, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.0016, Unit: schema.CostUnitToken, UnitSize: 1000},
	}
)
