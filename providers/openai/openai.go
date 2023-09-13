package openai

import (
	"github.com/kamushadenes/chloe/config"
	openai2 "github.com/sashabaranov/go-openai"
)

func init() {
	resetClient()
}

var openAIClient *openai2.Client

func resetClient() {
	openAIClient = openai2.NewClient(config.OpenAI.APIKey)
}
