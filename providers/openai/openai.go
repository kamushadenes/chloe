package openai

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/chat_models/openai"
	openai2 "github.com/sashabaranov/go-openai"
)

func init() {
	resetClient()
}

var completionClient common.Chat

var openAIClient *openai2.Client

func resetClient() {
	completionClient = openai.NewChatOpenAIWithDefaultModel(config.OpenAI.APIKey)

	openAIClient = openai2.NewClient(config.OpenAI.APIKey)
}
