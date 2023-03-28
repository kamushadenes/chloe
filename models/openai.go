package models

import "github.com/sashabaranov/go-openai"

type Model string

const (
	GPT35Turbo           Model = openai.GPT3Dot5Turbo
	GPT35Turbo0301       Model = openai.GPT3Dot5Turbo0301
	GPT4                 Model = openai.GPT4
	GPT40314             Model = openai.GPT40314
	GPT432K              Model = openai.GPT432K
	GPT432K0314          Model = openai.GPT432K0314
	DallE256X256         Model = "dall-e-256x256"
	DallE512X512         Model = "dall-e-512x512"
	DallE1024X1024       Model = "dall-e-1024x1024"
	Whisper1             Model = openai.Whisper1
	TextModerationStable Model = "text-moderation-stable"
	TextModerationLatest Model = "text-moderation-latest"
)

func ModelsToString(models ...Model) []string {
	var res []string
	for _, model := range models {
		res = append(res, string(model))
	}

	return res
}
