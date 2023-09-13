package openai

import (
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/sashabaranov/go-openai"
)

var (
	GPT35Turbo = &common.ChatModel{
		Name:             openai.GPT3Dot5Turbo,
		ContextSize:      4096,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		PromptCost:       &schema.CostObject{Price: 0.0015, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.002, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT35Turbo0301 = &common.ChatModel{
		Name:             openai.GPT3Dot5Turbo0301,
		ContextSize:      4096,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		PromptCost:       &schema.CostObject{Price: 0.0015, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.002, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT35Turbo0613 = &common.ChatModel{
		Name:             openai.GPT3Dot5Turbo0613,
		ContextSize:      4096,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		PromptCost:       &schema.CostObject{Price: 0.0015, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.002, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT35Turbo16K = &common.ChatModel{
		Name:             openai.GPT3Dot5Turbo16K,
		ContextSize:      16384,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		PromptCost:       &schema.CostObject{Price: 0.003, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.004, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT4 = &common.ChatModel{
		Name:             openai.GPT4,
		ContextSize:      8000,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.03, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.06, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT40314 = &common.ChatModel{
		Name:             openai.GPT40314,
		ContextSize:      8000,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.03, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.06, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT40613 = &common.ChatModel{
		Name:             openai.GPT40613,
		ContextSize:      8000,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.03, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.06, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT432K = &common.ChatModel{
		Name:             openai.GPT432K,
		ContextSize:      8000,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.06, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.12, Unit: schema.CostUnitToken, UnitSize: 1000},
	}

	GPT432K0314 = &common.ChatModel{
		Name:             openai.GPT432K0314,
		ContextSize:      8000,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.06, Unit: schema.CostUnitToken, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.12, Unit: schema.CostUnitToken, UnitSize: 1000},
	}
)
