package asr

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/audio_models/common"
	"github.com/kamushadenes/chloe/langchain/audio_models/openai"
	"github.com/kamushadenes/chloe/langchain/audio_models/whispercpp"
)

func NewASR(model *common.ASRModel) common.ASR {
	switch model {
	case openai.Whisper1:
		return openai.NewASROpenAI(config.OpenAI.APIKey, model)
	case whispercpp.Base, whispercpp.Tiny, whispercpp.Small, whispercpp.Medium, whispercpp.Large:
		return whispercpp.NewASRWhisperCpp(model)
	}

	return whispercpp.NewASRWhisperCpp(model)
}

func NewASRWithDefaultModel(provider config.STTProvider) common.ASR {
	switch provider {
	case config.OpenAISTT:
		return openai.NewASROpenAIWithDefaultModel(config.OpenAI.APIKey)
	case config.WhisperCppSTT:
		return whispercpp.NewASRWhisperCppWithDefaultModel()
	}

	return whispercpp.NewASRWhisperCppWithDefaultModel()
}
