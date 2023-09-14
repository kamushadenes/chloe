package base

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/chat_models/gpt4all"
	"github.com/kamushadenes/chloe/langchain/chat_models/openai"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func NewChat(model *common.ChatModel, user *memory.User) common.Chat {
	switch model {
	case openai.GPT35Turbo, openai.GPT35Turbo0301, openai.GPT4, openai.GPT40314, openai.GPT432K, openai.GPT432K0314:
		return openai.NewChatOpenAI(config.OpenAI.APIKey, model, user)
	}

	return gpt4all.NewChatGPT4All(model, user)
}

func NewChatWithDefaultModel(provider config.ChatProvider, user *memory.User) common.Chat {
	switch provider {
	case config.OpenAIChat:
		return openai.NewChatOpenAIWithDefaultModel(config.OpenAI.APIKey, user)
	case config.GPT4AllChat:
		return gpt4all.NewChatGPT4All(gpt4all.GetModel(config.GPT4All.DefaultModel), user)
	}

	return nil
}
