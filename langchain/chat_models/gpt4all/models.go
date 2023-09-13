package gpt4all

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
		UsageCost:        &schema.CostObject{Price: 0.002, Unit: schema.CostUnitToken, UnitSize: 1000},
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

	GPT4AllJV13Groovy = &common.ChatModel{
		Name:             "gpt4all-j-v1.3-groovy",
		ContextSize:      2048,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0, Unit: schema.CostUnitToken, UnitSize: 1000},
		Tokenizer:        "gpt-3.5-turbo",
	}

	MPT7BChat = &common.ChatModel{
		Name:             "mpt-7b-chat",
		ContextSize:      4096,
		ContextUnit:      schema.ContextUnitToken,
		TokensPerMessage: 3,
		TokensPerName:    1,
		UsageCost:        &schema.CostObject{Price: 0, Unit: schema.CostUnitToken, UnitSize: 1000},
		Tokenizer:        "gpt-3.5-turbo",
	}
)

func GetModel(model string) *common.ChatModel {
	switch model {
	case GPT35Turbo.Name:
		return GPT35Turbo
	case GPT4.Name:
		return GPT4
	case GPT4AllJV13Groovy.Name:
		return GPT4AllJV13Groovy
	case MPT7BChat.Name:
		return MPT7BChat
	}

	return nil
}
