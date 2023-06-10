package config

type ChatProvider string

const (
	OpenAIChat  ChatProvider = "openai"
	GPT4AllChat ChatProvider = "gpt4all"
)

type ChatConfig struct {
	Provider ChatProvider
}

var Chat = &ChatConfig{
	Provider: ChatProvider(envOrDefaultWithOptions("CHLOE_CHAT_PROVIDER", string(OpenAIChat),
		[]string{string(OpenAIChat), string(GPT4AllChat)})),
}
