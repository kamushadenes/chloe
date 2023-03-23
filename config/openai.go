package config

import (
	"github.com/sashabaranov/go-openai"
	"os"
	"time"
)

type ModelPurpose string
type ImageType string

type TimeoutType string

const (
	ModelPurposeCompletion     ModelPurpose = "completion"
	ModelPurposeTranscription  ModelPurpose = "transcription"
	ModelPurposeChainOfThought ModelPurpose = "chain_of_thought"
	ModelPurposeModeration                  = "moderation"

	ImageTypeGeneration ImageType = "generation"
	ImageTypeEdit       ImageType = "edit"
	ImageTypeVariation  ImageType = "variation"

	TimeoutTypeCompletion      TimeoutType = "completion"
	TimeoutTypeChainOfThought  TimeoutType = "chain_of_thought"
	TimeoutTypeTranscription   TimeoutType = "transcription"
	TimeoutTypeModeration      TimeoutType = "moderation"
	TimeoutTypeImageGeneration TimeoutType = "image_generation"
	TimeoutTypeImageEdit       TimeoutType = "image_edit"
	TimeoutTypeImageVariation  TimeoutType = "image_variation"
	TimeoutTypeSlowness        TimeoutType = "slowness"
)

type OpenAIConfig struct {
	MaxTokens                 map[string]int
	MinReplyTokens            map[string]int
	DefaultModel              map[ModelPurpose]string
	DefaultSize               map[ImageType]string
	APIKey                    string
	MessagesToKeepFullContent int
	ModerateMessages          bool
	Timeouts                  map[TimeoutType]time.Duration
}

var OpenAI = &OpenAIConfig{
	Timeouts: map[TimeoutType]time.Duration{
		TimeoutTypeCompletion:      60 * time.Second,
		TimeoutTypeChainOfThought:  60 * time.Second,
		TimeoutTypeTranscription:   60 * time.Second,
		TimeoutTypeModeration:      60 * time.Second,
		TimeoutTypeImageGeneration: 120 * time.Second,
		TimeoutTypeImageEdit:       120 * time.Second,
		TimeoutTypeImageVariation:  120 * time.Second,
		TimeoutTypeSlowness:        5 * time.Second,
	},
	ModerateMessages: true,
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
		ModelPurposeModeration:     "text-moderation-latest",
	},
	DefaultSize: map[ImageType]string{
		ImageTypeGeneration: "1024x1024",
		ImageTypeEdit:       "1024x1024",
		ImageTypeVariation:  "1024x1024",
	},
	APIKey:                    os.Getenv("OPENAI_API_KEY"),
	MessagesToKeepFullContent: 4,
}
