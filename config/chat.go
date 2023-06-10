package config

type ChatProvider string

const (
	OpenAIChat ChatProvider = "openai"
)

type ChatConfig struct {
	Provider ChatProvider
}

var Chat = &ChatConfig{
	Provider: Provider(envOrDefaultWithOptions("CHLOE_CHAT_PROVIDER", string(OpenAIChat),
		[]string{string(OpenAIChat)})),
}
