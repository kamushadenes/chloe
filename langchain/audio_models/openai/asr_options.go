package openai

import (
	"github.com/kamushadenes/chloe/langchain/audio_models/common"
	"github.com/sashabaranov/go-openai"
	"time"
)

type ASROptionsOpenAI struct {
	req     openai.AudioRequest
	timeout time.Duration
}

func NewASROptionsOpenAI() common.ASROptions {
	return &ASROptionsOpenAI{req: openai.AudioRequest{}}
}

func (c ASROptionsOpenAI) GetRequest() interface{} {
	return c.req
}

func (c ASROptionsOpenAI) GetAudioFile() string {
	return c.req.FilePath
}

func (c ASROptionsOpenAI) WithAudioFile(audioFile string) common.ASROptions {
	c.req.FilePath = audioFile
	return c
}

func (c ASROptionsOpenAI) WithModel(model string) common.ASROptions {
	c.req.Model = model
	return c
}

func (c ASROptionsOpenAI) WithPrompt(prompt string) common.ASROptions {
	c.req.Prompt = prompt
	return c
}

func (c ASROptionsOpenAI) WithTemperature(temperature float32) common.ASROptions {
	c.req.Temperature = temperature
	return c
}

func (c ASROptionsOpenAI) WithLanguage(language string) common.ASROptions {
	c.req.Language = language
	return c
}

func (c ASROptionsOpenAI) WithTimeout(timeout time.Duration) common.ASROptions {
	c.timeout = timeout
	return c
}

func (c ASROptionsOpenAI) GetTimeout() time.Duration {
	return c.timeout
}

func (c ASROptionsOpenAI) WithOutputFormat(format string) common.ASROptions {
	c.req.Format = openai.AudioResponseFormat(format)
	return c
}
