package common

import "time"

type TTSOptions interface {
	GetRequest() interface{}
	WithText(text string) TTSOptions
	WithVoice(voice string) TTSOptions
	WithAudioEncoding(encoding string) TTSOptions
	WithSpeakingRate(rate float64) TTSOptions
	WithPitch(pitch float64) TTSOptions
	WithVolumeGain(gain float64) TTSOptions
	WithLanguage(language string) TTSOptions
	WithTimeout(time.Duration) TTSOptions
	GetTimeout() time.Duration
	WithSimilarityBoost(float32) TTSOptions
	WithStability(float32) TTSOptions
	WithModel(TTSModel) TTSOptions
	WithPrompt(string) TTSOptions
}
