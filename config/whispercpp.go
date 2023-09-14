package config

type WhisperCppConfig struct {
	Model    string
	Language string
}

var WhisperCpp = &WhisperCppConfig{
	Model: envOrDefaultWithOptions("CHLOE_STT_WHISPERCPP_MODEL", "large",
		[]string{"base", "tiny", "small", "medium", "large"}),
	Language: envOrDefaultWithOptions("CHLOE_STT_WHISPERCPP_LANGUAGE", "pt",
		[]string{"en", "pt"}),
}
