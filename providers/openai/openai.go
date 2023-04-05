package openai

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/sashabaranov/go-openai"
)

func init() {
	resetClient()
}

var openAIClient *openai.Client

func resetClient() {
	openAIClient = openai.NewClient(config.OpenAI.APIKey)
}
