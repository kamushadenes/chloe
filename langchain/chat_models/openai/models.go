package openai

import (
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/sashabaranov/go-openai"
)


var (
	GPT35Turbo = &chat_models.ChatModel{
		Name:             openai.GPT3Dot5Turbo,
		ContextSize:      4096,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.002, Unit: schema.Token, UnitSize: 1000},
	}

	GPT35Turbo0301 = &chat_models.ChatModel{
		Name:             openai.GPT3Dot5Turbo0301,
		ContextSize:      4096,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &schema.CostObject{Price: 0.002, Unit: schema.Token, UnitSize: 1000},
	}

	GPT4 = &chat_models.ChatModel{
		Name:             openai.GPT4,
		ContextSize:      8000,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.03, Unit: schema.Token, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.06, Unit: schema.Token, UnitSize: 1000},
	}

	GPT40314 = &chat_models.ChatModel{
		Name:             openai.GPT40314,
		ContextSize:      8000,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.03, Unit: schema.Token, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.06, Unit: schema.Token, UnitSize: 1000},
	}

	GPT432K = &chat_models.ChatModel{
		Name:             openai.GPT432K,
		ContextSize:      8000,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.06, Unit: schema.Token, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.12, Unit: schema.Token, UnitSize: 1000},
	}

	GPT432K0314 = &chat_models.ChatModel{
		Name:             openai.GPT432K0314,
		ContextSize:      8000,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &schema.CostObject{Price: 0.06, Unit: schema.Token, UnitSize: 1000},
		CompletionCost:   &schema.CostObject{Price: 0.12, Unit: schema.Token, UnitSize: 1000},
	}
)
