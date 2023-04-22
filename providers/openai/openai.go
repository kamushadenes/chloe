package openai

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/chat_models/openai"
)

func init() {
	resetClient()
}

var completionClient common.Chat

func resetClient() {
	completionClient = openai.NewChatOpenAIWithDefaultModel(config.OpenAI.APIKey)
}
