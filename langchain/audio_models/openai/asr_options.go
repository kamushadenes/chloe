package openai

import (
	"github.com/kamushadenes/chloe/langchain/audio_models/common"
	"github.com/sashabaranov/go-openai"
	"time"
)

type ASROptionsGPT4All struct {
	req     openai.AudioRequest
	timeout time.Duration
}

func NewASROptionsGPT4All() common.ASROptions {
	return &ASROptionsGPT4All{req: openai.AudioRequest{}}
}

func (c ASROptionsGPT4All) GetRequest() interface{} {
	return c.req
}

func (c ASROptionsGPT4All) GetAudioFile() string {
	return c.req.FilePath
}

func (c ASROptionsGPT4All) WithAudioFile(audioFile string) common.ASROptions {
	c.req.FilePath = audioFile
	return c
}

func (c ASROptionsGPT4All) WithModel(model string) common.ASROptions {
	c.req.Model = model
	return c
}

func (c ASROptionsGPT4All) WithPrompt(prompt string) common.ASROptions {
	c.req.Prompt = prompt
	return c
}

func (c ASROptionsGPT4All) WithTemperature(temperature float32) common.ASROptions {
	c.req.Temperature = temperature
	return c
}

func (c ASROptionsGPT4All) WithLanguage(language string) common.ASROptions {
	c.req.Language = language
	return c
}

func (c ASROptionsGPT4All) WithTimeout(timeout time.Duration) common.ASROptions {
	c.timeout = timeout
	return c
}

func (c ASROptionsGPT4All) GetTimeout() time.Duration {
	return c.timeout
}

func (c ASROptionsGPT4All) WithOutputFormat(format string) common.ASROptions {
	c.req.Format = openai.AudioResponseFormat(format)
	return c
}
