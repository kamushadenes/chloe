package config

import "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"

type GCPConfig struct {
	TTSLanguageCode string
	TTSVoiceName    string
	TTSEncoding     texttospeechpb.AudioEncoding
}

var GCP = &GCPConfig{
	TTSLanguageCode: envOrDefault("CHLOE_TTS_LANGUAGE_CODE", "en-US"),
	TTSVoiceName:    envOrDefault("CHLOE_TTS_VOICE_NAME", "en-US-Wavenet-F"),
	TTSEncoding:     envOrDefaultGCPTTSEncoding("CHLOE_TTS_AUDIO_ENCONDING", texttospeechpb.AudioEncoding_MP3),
}
