package config

type TTSProvider string

const (
	GoogleTTS     TTSProvider = "google"
	ElevenLabsTTS TTSProvider = "elevenlabs"
)

type TTSConfig struct {
	Provider TTSProvider
}

var TTS = &TTSConfig{
	Provider: TTSProvider(envOrDefaultWithOptions("CHLOE_TTS_PROVIDER", string(GoogleTTS),
		[]string{string(GoogleTTS), string(ElevenLabsTTS)})),
}
