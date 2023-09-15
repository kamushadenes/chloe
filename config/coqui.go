package config

type CoquiConfig struct {
	TTSVoiceID string
	APIKey     string
	TTSSpeed   float64
	TTSModel   string
}

var Coqui = &CoquiConfig{
	TTSVoiceID: envOrDefault("CHLOE_TTS_COQUI_VOICE_ID", "default"),
	TTSSpeed:   envOrDefaultFloat64("CHLOE_TTS_COQUI_SPEED", 1),
	TTSModel: envOrDefaultWithOptions("CHLOE_TTS_COQUI_MODEL", "xtts",
		[]string{"coqui_v1", "xtts"}),
	APIKey: envOrDefault("CHLOE_TTS_COQUI_API_KEY", ""),
}
