package openai

import (
	"github.com/kamushadenes/chloe/config"
	openai "github.com/sashabaranov/go-openai"
)

var openAIClient = openai.NewClient(config.OpenAI.APIKey)
