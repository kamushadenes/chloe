package config

type ModelPurpose string
type ImagePurpose string

const (
	Completion    ModelPurpose = "completion"
	Transcription ModelPurpose = "transcription"
	Moderation    ModelPurpose = "moderation"
	Summarization ModelPurpose = "summarization"

	ImageGeneration ImagePurpose = "image_generation"
	ImageEdit       ImagePurpose = "image_edit"
	ImageVariation  ImagePurpose = "image_variation"
)

type OpenAIConfigModel struct {
	Completion    string
	Transcription string
	Moderation    string
	Summarization string
}

type OpenAIConfigImageSize struct {
	ImageGeneration string
	ImageEdit       string
	ImageVariation  string
}

type OpenAIConfig struct {
	MinReplyTokens            int
	DefaultModel              OpenAIConfigModel
	DefaultSize               OpenAIConfigImageSize
	APIKey                    string
	MessagesToKeepFullContent int
	ModerateMessages          bool
	SummarizeMessages         bool
	UseAzure                  bool
	AzureAPIVersion           string
	AzureBaseURL              string
	AzureEngine               string
	UseFunctions              bool
}

var OpenAI = &OpenAIConfig{
	ModerateMessages:  envOrDefaultBool("CHLOE_ENABLE_MESSAGE_MODERATION", true),
	SummarizeMessages: envOrDefaultBool("CHLOE_ENABLE_MESSAGE_SUMMARIZATION", true),

	MinReplyTokens: envOrDefaultInt("CHLOE_MIN_REPLY_TOKENS", 500),

	DefaultModel: OpenAIConfigModel{
		Completion:    envOrDefault("CHLOE_MODEL_COMPLETION", "gpt-3.5-turbo"),
		Transcription: envOrDefault("CHLOE_MODEL_TRANSCRIPTION", "whisper-1"),
		Moderation:    envOrDefault("CHLOE_MODEL_MODERATION", "text-moderation-latest"),
		Summarization: envOrDefault("CHLOE_MODEL_SUMMARIZATION", "gpt-3.5-turbo"),
	},

	DefaultSize: OpenAIConfigImageSize{
		ImageGeneration: envOrDefaultImageSize("CHLOE_IMAGE_GENERATION_SIZE", "1024x1024"),
		ImageEdit:       envOrDefaultImageSize("CHLOE_IMAGE_EDIT_SIZE", "1024x1024"),
		ImageVariation:  envOrDefaultImageSize("CHLOE_IMAGE_VARIATION_SIZE", "1024x1024"),
	},

	APIKey: envOrDefault("OPENAI_API_KEY", ""),

	MessagesToKeepFullContent: envOrDefaultInt("CHLOE_MESSAGES_TO_KEEP_FULL_CONTENT", 4),
	UseAzure:                  envOrDefaultBool("CHLOE_USE_AZURE", false),
	AzureAPIVersion:           envOrDefault("CHLOE_AZURE_API_VERSION", "2023-03-15-preview"),
	AzureBaseURL:              envOrDefault("CHLOE_AZURE_BASE_URL", ""),
	AzureEngine:               envOrDefault("CHLOE_AZURE_ENGINE", ""),
	UseFunctions:              envOrDefaultBool("CHLOE_USE_FUNCTIONS", true),
}
