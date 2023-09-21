package config

type LLMProvider string

const (
	OpenAILLM LLMProvider = "openai"
)

type LLMConfig struct {
	Provider LLMProvider
}

var LLM = &LLMConfig{
	Provider: LLMProvider(envOrDefaultWithOptions("CHLOE_LLM_PROVIDER", string(OpenAIChat),
		[]string{string(OpenAIChat)})),
}
