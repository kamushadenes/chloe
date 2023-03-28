package config

import (
	"github.com/kamushadenes/chloe/models"
)

type ModelPurpose string
type ImagePurpose string
type UnitType string

const (
	Completion     ModelPurpose = "completion"
	ChainOfThought ModelPurpose = "chain_of_thought"
	Transcription  ModelPurpose = "transcription"
	Moderation     ModelPurpose = "moderation"
	Summarization  ModelPurpose = "summarization"

	ImageGeneration ImagePurpose = "image_generation"
	ImageEdit       ImagePurpose = "image_edit"
	ImageVariation  ImagePurpose = "image_variation"

	Token  UnitType = "token"
	Image  UnitType = "image"
	Minute UnitType = "minute"
)

type OpenAIConfigModel struct {
	Completion     models.Model
	Transcription  models.Model
	ChainOfThought models.Model
	Moderation     models.Model
	Summarization  models.Model
}

type OpenAIConfigImageSize struct {
	ImageGeneration string
	ImageEdit       string
	ImageVariation  string
}

type OpenAICostObject struct {
	Price    float64
	Unit     UnitType
	UnitSize int
}

type OpenAICostModel struct {
	Name       string
	Usage      *OpenAICostObject
	Prompt     *OpenAICostObject
	Completion *OpenAICostObject
}

type OpenAIConfig struct {
	MaxTokens                 map[models.Model]int
	MinReplyTokens            int
	DefaultModel              OpenAIConfigModel
	DefaultSize               OpenAIConfigImageSize
	APIKey                    string
	MessagesToKeepFullContent int
	ModerateMessages          bool
	Costs                     map[models.Model]*OpenAICostModel
}

var OpenAI = &OpenAIConfig{
	ModerateMessages: envOrDefaultBool("CHLOE_ENABLE_MESSAGE_MODERATION", true),

	MaxTokens: map[models.Model]int{
		models.GPT35Turbo: envOrDefaultInt("CHLOE_MAX_TOKENS_GPT35Turbo", 4096),
		models.GPT4:       envOrDefaultInt("CHLOE_MAX_TOKENS_GPT4", 8000),
		models.GPT432K:    envOrDefaultInt("CHLOE_MAX_TOKENS_GPT432K", 32000),
	},

	MinReplyTokens: envOrDefaultInt("CHLOE_MIN_REPLY_TOKENS", 500),

	DefaultModel: OpenAIConfigModel{
		Completion:     envOrDefaultCompletionModel("CHLOE_MODEL_COMPLETION", models.GPT35Turbo),
		ChainOfThought: envOrDefaultCompletionModel("CHLOE_MODEL_CHAIN_OF_THOUGHT", models.GPT35Turbo),
		Transcription:  envOrDefaultTranscriptionModel("CHLOE_MODEL_TRANSCRIPTION", models.Whisper1),
		Moderation:     envOrDefaultModerationModel("CHLOE_MODEL_MODERATION", models.TextModerationLatest),
		Summarization:  envOrDefaultCompletionModel("CHLOE_MODEL_SUMMARIZATION", models.GPT35Turbo),
	},

	DefaultSize: OpenAIConfigImageSize{
		ImageGeneration: envOrDefaultImageSize("CHLOE_IMAGE_GENERATION_SIZE", "1024x1024"),
		ImageEdit:       envOrDefaultImageSize("CHLOE_IMAGE_EDIT_SIZE", "1024x1024"),
		ImageVariation:  envOrDefaultImageSize("CHLOE_IMAGE_VARIATION_SIZE", "1024x1024"),
	},
	APIKey: mustEnv("OPENAI_API_KEY"),

	MessagesToKeepFullContent: envOrDefaultInt("CHLOE_MESSAGES_TO_KEEP_FULL_CONTENT", 4),

	Costs: map[models.Model]*OpenAICostModel{
		models.GPT35Turbo: {
			Name:  "GPT-3.5 Turbo",
			Usage: &OpenAICostObject{Price: 0.002, Unit: Token, UnitSize: 1000},
		},
		models.GPT4: {
			Name:       "GPT-4",
			Prompt:     &OpenAICostObject{Price: 0.03, Unit: Token, UnitSize: 1000},
			Completion: &OpenAICostObject{Price: 0.06, Unit: Token, UnitSize: 1000},
		},
		models.GPT432K: {
			Name:       "GPT-4 32K",
			Prompt:     &OpenAICostObject{Price: 0.06, Unit: Token, UnitSize: 1000},
			Completion: &OpenAICostObject{Price: 0.12, Unit: Token, UnitSize: 1000},
		},
		models.DallE256X256: {
			Name:  "DALL-E 256x256",
			Usage: &OpenAICostObject{Price: 0.016, Unit: Image, UnitSize: 1},
		},
		models.DallE512X512: {
			Name:  "DALL-E 512x512",
			Usage: &OpenAICostObject{Price: 0.018, Unit: Image, UnitSize: 1},
		},
		models.DallE1024X1024: {
			Name:  "DALL-E 1024x1024",
			Usage: &OpenAICostObject{Price: 0.020, Unit: Image, UnitSize: 1},
		},
		models.Whisper1: {
			Name:  "Whisper 1",
			Usage: &OpenAICostObject{Price: 0.006, Unit: Minute, UnitSize: 1},
		},
	},
}
