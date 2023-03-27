package config

import (
	"github.com/sashabaranov/go-openai"
)

type OpenAIConfigModel struct {
	Completion     string
	Transcription  string
	ChainOfThought string
	Moderation     string
	Summarization  string
}

type OpenAIConfigImageSize struct {
	ImageGeneration string
	ImageEdit       string
	ImageVariation  string
}

type OpenAIConfig struct {
	MaxTokens                 map[string]int
	MinReplyTokens            map[string]int
	DefaultModel              OpenAIConfigModel
	DefaultSize               OpenAIConfigImageSize
	APIKey                    string
	MessagesToKeepFullContent int
	ModerateMessages          bool
}

var OpenAI = &OpenAIConfig{
	ModerateMessages: envOrDefaultBool("CHLOE_ENABLE_MESSAGE_MODERATION", true),
	MaxTokens: map[string]int{
		openai.GPT3Dot5Turbo: envOrDefaultInt("CHLOE_MAX_TOKENS_GPT3Dot5Turbo", 4096),
	},
	MinReplyTokens: map[string]int{
		openai.GPT3Dot5Turbo: envOrDefaultInt("CHLOE_MIN_REPLY_TOKENS_GPT3Dot5Turbo", 500),
	},
	DefaultModel: OpenAIConfigModel{
		Completion:     envOrDefault("CHLOE_MODEL_COMPLETION", openai.GPT3Dot5Turbo),
		ChainOfThought: envOrDefault("CHLOE_MODEL_CHAIN_OF_THOUGHT", openai.GPT3Dot5Turbo),
		Transcription:  envOrDefaultWithOptions("CHLOE_MODEL_TRANSCRIPTION", openai.Whisper1, []string{openai.Whisper1}),
		Moderation:     envOrDefaultWithOptions("CHLOE_MODEL_MODERATION", "text-moderation-latest", []string{"text-moderation-latest", "text-moderation-stable"}),
		Summarization:  envOrDefault("CHLOE_MODEL_SUMMARIZATION", openai.GPT3Dot5Turbo),
	},
	DefaultSize: OpenAIConfigImageSize{
		ImageGeneration: envOrDefaultWithOptions("CHLOE_IMAGE_GENERATION_SIZE", "1024x1024", []string{"1024x1024", "512x512", "256x256"}),
		ImageEdit:       envOrDefaultWithOptions("CHLOE_IMAGE_EDIT_SIZE", "1024x1024", []string{"1024x1024", "512x512", "256x256"}),
		ImageVariation:  envOrDefaultWithOptions("CHLOE_IMAGE_VARIATION_SIZE", "1024x1024", []string{"1024x1024", "512x512", "256x256"}),
	},
	APIKey:                    mustEnv("OPENAI_API_KEY"),
	MessagesToKeepFullContent: envOrDefaultInt("CHLOE_MESSAGES_TO_KEEP_FULL_CONTENT", 4),
}
