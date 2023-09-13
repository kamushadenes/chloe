package config

import "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"

type GoogleConfig struct {
	TTSLanguageCode string
	TTSVoiceName    string
	TTSEncoding     texttospeechpb.AudioEncoding
	TTSSpeakingRate float64
	TTSPitch        float64
	TTSVolumeGain   float64
}

var Google = &GoogleConfig{
	TTSLanguageCode: envOrDefault("CHLOE_TTS_GOOGLE_LANGUAGE_CODE", "en-US"),
	TTSVoiceName:    envOrDefault("CHLOE_TTS_GOOGLE_VOICE_NAME", "en-US-Wavenet-F"),
	TTSEncoding:     envOrDefaultGCPTTSEncoding("CHLOE_TTS_GOOGLE_AUDIO_ENCONDING", texttospeechpb.AudioEncoding_MP3),
	TTSSpeakingRate: envOrDefaultFloat64InRange("CHLOE_TTS_GOOGLE_SPEAKING_RATE", 1, 0.25, 4),
	TTSPitch:        envOrDefaultFloat64InRange("CHLOE_TTS_GOOGLE_PITCH", 0, -20, 20),
	TTSVolumeGain:   envOrDefaultFloat64InRange("CHLOE_TTS_GOOGLE_VOLUME_GAIN", 0, -96, 16),
}
