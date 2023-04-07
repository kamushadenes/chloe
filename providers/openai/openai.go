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
	if config.OpenAI.UseAzure {
		openAIClient = openai.NewClientWithConfig(openai.DefaultAzureConfig(config.OpenAI.APIKey, config.OpenAI.AzureBaseURL, config.OpenAI.AzureEngine))
	} else {
		openAIClient = openai.NewClient(config.OpenAI.APIKey)
	}
}
