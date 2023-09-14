package base

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/langchain/tts/elevenlabs"
	"github.com/kamushadenes/chloe/langchain/tts/google"
)

func NewTTS(model common.TTSModel) common.TTS {
	switch model {
	case google.GoogleTTS:
		return google.NewTTSGoogle()
	}

	return google.NewTTSGoogle()
}

func NewTTSWithDefaultModel(provider config.TTSProvider) common.TTS {
	switch provider {
	case config.GoogleTTS:
		return google.NewTTSGoogle()
	case config.ElevenLabsTTS:
		return elevenlabs.NewTTSElevenLabs()
	}

	return google.NewTTSGoogle()
}
