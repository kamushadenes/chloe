package config

import (
	"github.com/sashabaranov/go-openai"
)

type ModelPurpose string
type ImagePurpose string

const (
	Completion     ModelPurpose = "completion"
	ChainOfThought ModelPurpose = "chain_of_thought"
	Transcription  ModelPurpose = "transcription"
	Moderation     ModelPurpose = "moderation"
	Summarization  ModelPurpose = "summarization"

	ImageGeneration ImagePurpose = "image_generation"
	ImageEdit       ImagePurpose = "image_edit"
	ImageVariation  ImagePurpose = "image_variation"
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
	MinReplyTokens            int
	DefaultModel              OpenAIConfigModel
	DefaultSize               OpenAIConfigImageSize
	APIKey                    string
	MessagesToKeepFullContent int
	ModerateMessages          bool
}

func (c *OpenAIConfig) GetMaxTokens(model string) int {
	switch model {
	case openai.GPT3Dot5Turbo, openai.GPT3Dot5Turbo0301:
		model = openai.GPT3Dot5Turbo
	case openai.GPT4, openai.GPT40314:
		model = openai.GPT4
	case openai.GPT432K, openai.GPT432K0314:
		model = openai.GPT432K
	}

	return c.MaxTokens[model]
}

func (c *OpenAIConfig) GetMinReplyTokens() int {
	return c.MinReplyTokens
}

func (c *OpenAIConfig) GetTokenizerModel(model string) string {
	switch model {
	case openai.GPT3Dot5Turbo, openai.GPT3Dot5Turbo0301:
		return openai.GPT3Dot5Turbo0301
	case openai.GPT4, openai.GPT40314:
		return openai.GPT40314
	case openai.GPT432K, openai.GPT432K0314:
		return openai.GPT432K0314
	default:
		return model
	}
}

func (c *OpenAIConfig) GetModel(purpose ModelPurpose) string {
	switch purpose {
	case Completion:
		return c.DefaultModel.Completion
	case ChainOfThought:
		return c.DefaultModel.ChainOfThought
	case Transcription:
		return c.DefaultModel.Transcription
	case Moderation:
		return c.DefaultModel.Moderation
	case Summarization:
		return c.DefaultModel.Summarization
	default:
		return c.DefaultModel.Completion
	}
}

func (c *OpenAIConfig) GetImageSize(purpose ImagePurpose) string {
	switch purpose {
	case ImageGeneration:
		return c.DefaultSize.ImageGeneration
	case ImageEdit:
		return c.DefaultSize.ImageEdit
	case ImageVariation:
		return c.DefaultSize.ImageVariation
	default:
		return c.DefaultSize.ImageGeneration
	}
}

var OpenAI = &OpenAIConfig{
	ModerateMessages: envOrDefaultBool("CHLOE_ENABLE_MESSAGE_MODERATION", true),
	MaxTokens: map[string]int{
		openai.GPT3Dot5Turbo: envOrDefaultInt("CHLOE_MAX_TOKENS_GPT3Dot5Turbo", 4096),
		openai.GPT4:          envOrDefaultInt("CHLOE_MAX_TOKENS_GPT4", 8000),
		openai.GPT432K:       envOrDefaultInt("CHLOE_MAX_TOKENS_GPT432K", 32000),
	},
	MinReplyTokens: envOrDefaultInt("CHLOE_MIN_REPLY_TOKENS", 500),
	DefaultModel: OpenAIConfigModel{
		Completion:     envOrDefaultCompletionModel("CHLOE_MODEL_COMPLETION", openai.GPT3Dot5Turbo),
		ChainOfThought: envOrDefaultCompletionModel("CHLOE_MODEL_CHAIN_OF_THOUGHT", openai.GPT3Dot5Turbo),
		Transcription:  envOrDefaultTranscriptionModel("CHLOE_MODEL_TRANSCRIPTION", openai.Whisper1),
		Moderation:     envOrDefaultModerationModel("CHLOE_MODEL_MODERATION", "text-moderation-latest"),
		Summarization:  envOrDefaultCompletionModel("CHLOE_MODEL_SUMMARIZATION", openai.GPT3Dot5Turbo),
	},
	DefaultSize: OpenAIConfigImageSize{
		ImageGeneration: envOrDefaultImageSize("CHLOE_IMAGE_GENERATION_SIZE", "1024x1024"),
		ImageEdit:       envOrDefaultImageSize("CHLOE_IMAGE_EDIT_SIZE", "1024x1024"),
		ImageVariation:  envOrDefaultImageSize("CHLOE_IMAGE_VARIATION_SIZE", "1024x1024"),
	},
	APIKey:                    mustEnv("OPENAI_API_KEY"),
	MessagesToKeepFullContent: envOrDefaultInt("CHLOE_MESSAGES_TO_KEEP_FULL_CONTENT", 4),
}
