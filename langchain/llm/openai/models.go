package openai

import (
	"github.com/kamushadenes/chloe/langchain/llm"
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/sashabaranov/go-openai"
)

var (
	Ada = &llm.LLMModel{
		Name:             openai.GPT3Ada,
		ContextSize:      2049,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0004, Unit: schema.Token, UnitSize: 1000},
	}

	TextAda001 = &llm.LLMModel{
		Name:             openai.GPT3TextAda001,
		ContextSize:      2049,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0004, Unit: schema.Token, UnitSize: 1000},
	}

	Babbage = &llm.LLMModel{
		Name:             openai.GPT3Babbage,
		ContextSize:      2049,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0005, Unit: schema.Token, UnitSize: 1000},
	}

	TextBabbage001 = &llm.LLMModel{
		Name:             openai.GPT3TextBabbage001,
		ContextSize:      2049,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.0005, Unit: schema.Token, UnitSize: 1000},
	}

	Curie = &llm.LLMModel{
		Name:             openai.GPT3Curie,
		ContextSize:      2049,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0020, Unit: schema.Token, UnitSize: 1000},
	}

	TextCurie001 = &llm.LLMModel{
		Name:             openai.GPT3TextCurie001,
		ContextSize:      2049,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0020, Unit: schema.Token, UnitSize: 1000},
	}

	Davinci = &llm.LLMModel{
		Name:             openai.GPT3Davinci,
		ContextSize:      2049,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0200, Unit: schema.Token, UnitSize: 1000},
	}

	TextDavinci002 = &llm.LLMModel{
		Name:             openai.GPT3TextDavinci002,
		ContextSize:      4097,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0200, Unit: schema.Token, UnitSize: 1000},
	}

	TextDavinci003 = &llm.LLMModel{
		Name:             openai.GPT3TextDavinci003,
		ContextSize:      4097,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0.0200, Unit: schema.Token, UnitSize: 1000},
	}
)