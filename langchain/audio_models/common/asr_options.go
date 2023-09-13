package common

import "time"

type ASROptions interface {
	GetRequest() interface{}
	GetAudioFile() string
	WithAudioFile(audioFile string) ASROptions
	WithModel(model string) ASROptions
	WithPrompt(prompt string) ASROptions
	WithTemperature(temperature float32) ASROptions
	WithLanguage(language string) ASROptions
	WithTimeout(time.Duration) ASROptions
	GetTimeout() time.Duration
	WithOutputFormat(format string) ASROptions
}
