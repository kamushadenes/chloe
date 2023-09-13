package base

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/diffusion_models/common"
	"github.com/kamushadenes/chloe/langchain/diffusion_models/openai"
)

func NewDiffusion(model *common.DiffusionModel) common.Diffusion {
	switch model {
	case openai.DallE256X256, openai.DallE512X512, openai.DallE1024X1024:
		return openai.NewDiffusionOpenAI(config.OpenAI.APIKey)
	}

	return nil
}

func NewDiffusionWithDefaultModel(provider config.DiffusionProvider) common.Diffusion {
	switch provider {
	case config.OpenAIDiffusion:
		return openai.NewDiffusionOpenAI(config.OpenAI.APIKey)
	}

	return nil
}
