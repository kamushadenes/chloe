package models

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/tokenizer"
	"github.com/kamushadenes/chloe/utils"
	"github.com/sashabaranov/go-openai"
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

	panic(errors.Wrap(errors.ErrModelNotFound, fmt.Errorf("model %s not found", name)))
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
		totalCost += m.UsageCost.Price * float64(m.CountChatCompletionTokens(messages)) / float64(m.UsageCost.UnitSize)
		totalCost += m.UsageCost.Price * (float64(tokenizer.CountTokens(m.String(), response)) / float64(m.UsageCost.UnitSize))
	} else if m.CompletionCost != nil && m.PromptCost != nil {
		totalCost += m.PromptCost.Price * float64(m.CountChatCompletionTokens(messages)) / float64(m.PromptCost.UnitSize)
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
