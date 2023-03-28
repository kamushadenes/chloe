package config

import "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"

type GCPConfig struct {
	TTSLanguageCode string
	TTSVoiceName    string
	TTSEncoding     texttospeechpb.AudioEncoding
	TTSSpeakingRate float64
	TTSPitch        float64
	TTSVolumeGain   float64
}

var GCP = &GCPConfig{
	TTSLanguageCode: envOrDefault("CHLOE_TTS_LANGUAGE_CODE", "en-US"),
	TTSVoiceName:    envOrDefault("CHLOE_TTS_VOICE_NAME", "en-US-Wavenet-F"),
	TTSEncoding:     envOrDefaultGCPTTSEncoding("CHLOE_TTS_AUDIO_ENCONDING", texttospeechpb.AudioEncoding_MP3),
	TTSSpeakingRate: envOrDefaultFloat64InRange("CHLOE_TTS_SPEAKING_RATE", 1, 0.25, 4),
	TTSPitch:        envOrDefaultFloat64InRange("CHLOE_TTS_PITCH", 0, -20, 20),
	TTSVolumeGain:   envOrDefaultFloat64InRange("CHLOE_TTS_VOLUME_GAIN", 0, -96, 16),
}
