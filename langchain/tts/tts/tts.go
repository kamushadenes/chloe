package base

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/langchain/tts/coqui"
	"github.com/kamushadenes/chloe/langchain/tts/elevenlabs"
	"github.com/kamushadenes/chloe/langchain/tts/google"
)

func NewTTS(model common.TTSModel) common.TTS {
	switch model {
	case google.GoogleTTS:
		return google.NewTTSGoogle()
	case elevenlabs.ElevenLabsTTS:
		return elevenlabs.NewTTSElevenLabs()
	case coqui.CoquiV1TTS, coqui.XTTS:
		return coqui.NewTTSCoqui(config.Coqui.APIKey)
	}

	return google.NewTTSGoogle()
}

func NewTTSWithDefaultModel(provider config.TTSProvider) common.TTS {
	switch provider {
	case config.GoogleTTS:
		return google.NewTTSGoogle()
	case config.ElevenLabsTTS:
		return elevenlabs.NewTTSElevenLabs()
	case config.CoquiTTS:
		return coqui.NewTTSCoqui(config.Coqui.APIKey)
	}

	return google.NewTTSGoogle()
}
