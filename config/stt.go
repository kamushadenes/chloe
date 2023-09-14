package config

type STTProvider string

const (
	WhisperCppSTT STTProvider = "whispercpp"
	OpenAISTT     STTProvider = "openai"
)

type STTConfig struct {
	Provider STTProvider
}

var STT = &STTConfig{
	Provider: STTProvider(envOrDefaultWithOptions("CHLOE_STT_PROVIDER", string(WhisperCppSTT),
		[]string{string(WhisperCppSTT), string(OpenAISTT)})),
}
