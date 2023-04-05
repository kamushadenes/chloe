package models

import (
	"github.com/sashabaranov/go-openai"
)

type UnitType string

const (
	Token  UnitType = "token"
	Image  UnitType = "image"
	Minute UnitType = "minute"
)

type OpenAICostObject struct {
	Price    float64
	Unit     UnitType
	UnitSize int
}

type Model struct {
	Name             string
	ContextSize      int
	TokensPerMessage int
	TokensPerName    int
	UsageCost        *OpenAICostObject
	PromptCost       *OpenAICostObject
	CompletionCost   *OpenAICostObject
}

var (
	GPT35Turbo = &Model{
		Name:             openai.GPT3Dot5Turbo,
		ContextSize:      4096,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &OpenAICostObject{Price: 0.002, Unit: Token, UnitSize: 1000},
	}

	GPT35Turbo0301 = &Model{
		Name:             openai.GPT3Dot5Turbo0301,
		ContextSize:      4096,
		TokensPerMessage: 4,
		TokensPerName:    -1,
		UsageCost:        &OpenAICostObject{Price: 0.002, Unit: Token, UnitSize: 1000},
	}

	GPT4 = &Model{
		Name:             openai.GPT4,
		ContextSize:      8000,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &OpenAICostObject{Price: 0.03, Unit: Token, UnitSize: 1000},
		CompletionCost:   &OpenAICostObject{Price: 0.06, Unit: Token, UnitSize: 1000},
	}

	GPT40314 = &Model{
		Name:             openai.GPT40314,
		ContextSize:      8000,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &OpenAICostObject{Price: 0.03, Unit: Token, UnitSize: 1000},
		CompletionCost:   &OpenAICostObject{Price: 0.06, Unit: Token, UnitSize: 1000},
	}

	GPT432K = &Model{
		Name:             openai.GPT432K,
		ContextSize:      8000,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &OpenAICostObject{Price: 0.06, Unit: Token, UnitSize: 1000},
		CompletionCost:   &OpenAICostObject{Price: 0.12, Unit: Token, UnitSize: 1000},
	}

	GPT432K0314 = &Model{
		Name:             openai.GPT432K0314,
		ContextSize:      8000,
		TokensPerMessage: 3,
		TokensPerName:    1,
		PromptCost:       &OpenAICostObject{Price: 0.06, Unit: Token, UnitSize: 1000},
		CompletionCost:   &OpenAICostObject{Price: 0.12, Unit: Token, UnitSize: 1000},
	}

	DallE256X256 = &Model{
		Name:      "dall-e-256x256",
		UsageCost: &OpenAICostObject{Price: 0.016, Unit: Image, UnitSize: 1},
	}

	DallE512X512 = &Model{
		Name:      "dall-e-512x512",
		UsageCost: &OpenAICostObject{Price: 0.018, Unit: Image, UnitSize: 1},
	}

	DallE1024X1024 = &Model{
		Name:      "dall-e-1024x1024",
		UsageCost: &OpenAICostObject{Price: 0.020, Unit: Image, UnitSize: 1},
	}

	Whisper1 = &Model{
		Name:      openai.Whisper1,
		UsageCost: &OpenAICostObject{Price: 0.006, Unit: Minute, UnitSize: 1},
	}

	TextModerationStable = &Model{
		Name: "text-moderation-stable",
	}

	TextModerationLatest = &Model{
		Name: "text-moderation-latest",
	}
)
