package models

import (
	"github.com/kamushadenes/chloe/tokenizer"
	"github.com/kamushadenes/chloe/utils"
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

func GetModel(name string) *Model {
	switch name {
	case GPT35Turbo.Name:
		return GPT35Turbo
	case GPT35Turbo0301.Name:
		return GPT35Turbo0301
	case GPT4.Name:
		return GPT4
	case GPT40314.Name:
		return GPT40314
	case GPT432K.Name:
		return GPT432K
	case GPT432K0314.Name:
		return GPT432K0314
	case DallE256X256.Name:
		return DallE256X256
	case DallE512X512.Name:
		return DallE512X512
	case DallE1024X1024.Name:
		return DallE1024X1024
	case Whisper1.Name:
		return Whisper1
	case TextModerationStable.Name:
		return TextModerationStable
	case TextModerationLatest.Name:
		return TextModerationLatest
	}

	return nil
}

func ModelsToString(models ...*Model) []string {
	var res []string
	for _, model := range models {
		res = append(res, model.String())
	}

	return res
}

func (m Model) String() string {
	return m.Name
}

func (m Model) GetContextSize() int {
	return m.ContextSize
}

func (m Model) CountTokens(text string) int {
	return tokenizer.CountTokens(m.String(), text)
}

func (m Model) CountChatCompletionTokens(messages []openai.ChatCompletionMessage) int {
	var tokens int

	for k := range messages {
		tokens += m.TokensPerMessage

		tokens += tokenizer.CountTokens(m.String(), messages[k].Role)
		tokens += tokenizer.CountTokens(m.String(), messages[k].Content)
		tokens += tokenizer.CountTokens(m.String(), messages[k].Name)
		if messages[k].Name != "" && messages[k].Role == "" {
			tokens -= m.TokensPerName // if there's a name, the role can be ommited, so we need to remove one token if it's empty
		}
	}

	tokens += 3 // every reply is primed with <im_start>assistant

	return tokens
}

func (m Model) GetCostForTokens(count int) float64 {
	if m.UsageCost != nil {
		return utils.RoundFloat(m.UsageCost.Price*float64(count), 5)
	} else if m.CompletionCost != nil {
		return utils.RoundFloat(m.CompletionCost.Price*float64(count), 5)
	}

	return 0
}

func (m Model) GetChatCompletionCost(messages []openai.ChatCompletionMessage, response string) float64 {
	var totalCost float64

	if m.UsageCost != nil {
		for k := range messages {
			totalCost += m.UsageCost.Price * (float64(tokenizer.CountTokens(m.String(), messages[k].Content)) / float64(m.UsageCost.UnitSize))
		}
		totalCost += m.UsageCost.Price * (float64(tokenizer.CountTokens(m.String(), response)) / float64(m.UsageCost.UnitSize))
	} else if m.CompletionCost != nil && m.PromptCost != nil {
		for k := range messages {
			totalCost += m.PromptCost.Price * (float64(tokenizer.CountTokens(m.String(), messages[k].Content)) / float64(m.PromptCost.UnitSize))
		}
		totalCost += m.CompletionCost.Price * (float64(tokenizer.CountTokens(m.String(), response)) / float64(m.CompletionCost.UnitSize))
	}

	return utils.RoundFloat(totalCost, 5)
}

func (m Model) GetGenerationCost(count int) float64 {
	if m.UsageCost != nil {
		return utils.RoundFloat(m.UsageCost.Price*float64(count), 5)
	}

	return 0
}

func (m Model) GetTranscriptionCost(minutes int) float64 {
	if m.UsageCost != nil {
		return utils.RoundFloat(m.UsageCost.Price*float64(minutes), 5)
	}

	return 0
}
