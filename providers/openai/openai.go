package openai

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/sashabaranov/go-openai"
)

var openAIClient *openai.Client

func resetClient() {
	openAIClient = openai.NewClient(config.OpenAI.APIKey)
}

func init() {
	resetClient()
}
