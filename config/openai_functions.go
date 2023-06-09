package config

import (
	"github.com/kamushadenes/chloe/models"
)

func (c *OpenAIConfig) GetMinReplyTokens() int {
	return c.MinReplyTokens
}

func (c *OpenAIConfig) GetTokenizerModel(model *models.Model) *models.Model {
	switch model {
	case models.GPT35Turbo, models.GPT35Turbo0301:
		return models.GPT35Turbo0301
	case models.GPT4, models.GPT40314:
		return models.GPT40314
	case models.GPT432K, models.GPT432K0314:
		return models.GPT432K0314
	default:
		return model
	}
}

func (c *OpenAIConfig) GetModel(purpose ModelPurpose) *models.Model {
	switch purpose {
	case Completion:
		return c.DefaultModel.Completion
	case ChainOfThought:
		return c.DefaultModel.ChainOfThought
	case Transcription:
		return c.DefaultModel.Transcription
	case Moderation:
		return c.DefaultModel.Moderation
	case Summarization:
		return c.DefaultModel.Summarization
	default:
		return c.DefaultModel.Completion
	}
}

func (c *OpenAIConfig) GetImageSize(purpose ImagePurpose) string {
	switch purpose {
	case ImageGeneration:
		return c.DefaultSize.ImageGeneration
	case ImageEdit:
		return c.DefaultSize.ImageEdit
	case ImageVariation:
		return c.DefaultSize.ImageVariation
	default:
		return c.DefaultSize.ImageGeneration
	}
}

func (c *OpenAIConfig) GetModelCostInfo(purpose ModelPurpose) (promptPrice float64, promptUnitSize int, completionPrice float64, completionUnitSize int) {
	model := c.GetModel(purpose)

	if model.UsageCost != nil {
		promptPrice = model.UsageCost.Price
		completionPrice = model.UsageCost.Price
		promptUnitSize = model.UsageCost.UnitSize
		completionUnitSize = model.UsageCost.UnitSize
	} else {
		promptPrice = model.PromptCost.Price
		completionPrice = model.CompletionCost.Price
		promptUnitSize = model.PromptCost.UnitSize
		completionUnitSize = model.CompletionCost.UnitSize
	}

	return
}
