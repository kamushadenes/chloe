package config

import (
	"github.com/sashabaranov/go-openai"
	"os"
)

type ModelPurpose string
type ImageType string

const (
	ModelPurposeCompletion     ModelPurpose = "completion"
	ModelPurposeTranscription  ModelPurpose = "transcription"
	ModelPurposeChainOfThought ModelPurpose = "chain_of_thought"

	ImageTypeGeneration ImageType = "generation"
	ImageTypeEdit       ImageType = "edit"
	ImageTypeVariation  ImageType = "variation"
)

type OpenAIConfig struct {
	MaxTokens      map[string]int
	MinReplyTokens map[string]int
	DefaultModel   map[ModelPurpose]string
	DefaultSize    map[ImageType]string
	APIKey         string
}

var OpenAI = &OpenAIConfig{
	MaxTokens: map[string]int{
		openai.GPT3Dot5Turbo: 4096,
	},
	MinReplyTokens: map[string]int{
		openai.GPT3Dot5Turbo: 500,
	},
	DefaultModel: map[ModelPurpose]string{
		ModelPurposeCompletion:     openai.GPT3Dot5Turbo,
		ModelPurposeChainOfThought: openai.GPT3Dot5Turbo,
		ModelPurposeTranscription:  openai.Whisper1,
	},
	DefaultSize: map[ImageType]string{
		ImageTypeGeneration: "1024x1024",
		ImageTypeEdit:       "1024x1024",
		ImageTypeVariation:  "1024x1024",
	},
	APIKey: os.Getenv("OPENAI_API_KEY"),
}
