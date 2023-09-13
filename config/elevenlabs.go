package config

type ElevenLabsConfig struct {
	TTSLanguageCode    string
	TTSVoiceName       string
	TTSStability       float32
	TTSSimilarityBoost float32
	APIKey             string
}

var ElevenLabs = &ElevenLabsConfig{
	TTSLanguageCode:    envOrDefault("CHLOE_TTS_ELEVENLABS_LANGUAGE_CODE", "eleven_multilingual_v2"),
	TTSVoiceName:       envOrDefault("CHLOE_TTS_ELEVENLABS_VOICE_NAME", "default"),
	TTSStability:       envOrDefaultFloat32InRange("CHLOE_TTS_ELEVENLABS_STABILITY", 60, 0, 100),
	TTSSimilarityBoost: envOrDefaultFloat32InRange("CHLOE_TTS_ELEVENLABS_SIMILARITY_BOOST", 90, 0, 100),
	APIKey:             envOrDefault("CHLOE_TTS_ELEVENLABS_API_KEY", ""),
}
