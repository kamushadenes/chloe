package base

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/llm/common"
	"github.com/kamushadenes/chloe/langchain/llm/openai"
)

func NewLLM(model *common.LLMModel) common.LLM {
	switch model {
	case openai.GPT35Turbo, openai.GPT4, openai.GPT432K, openai.Babbage, openai.Davinci, openai.GPT3Dot5TurboInstruct:
		return openai.NewLLMOpenAI(config.OpenAI.APIKey, model)
	}

	return openai.NewLLMOpenAI(config.OpenAI.APIKey, model)
}

func NewLLMWithDefaultModel(provider config.LLMProvider) common.LLM {
	switch provider {
	case config.OpenAILLM:
		return openai.NewLLMOpenAIWithDefaultModel(config.OpenAI.APIKey)
	}

	return nil
}
