package coqui

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/common"
)

type TTSOptionsCoqui struct {
	timeout  time.Duration
	voice    string
	text     string
	urlPath  string
	prompt   string
	language string
}

func NewTTSOptionsCoqui() common.TTSOptions {
	return &TTSOptionsCoqui{
		voice: config.Coqui.TTSVoiceID,
	}
}

func (c TTSOptionsCoqui) GetRequest() interface{} {
	u, _ := url.Parse("https://app.coqui.ai")
	u.Path = c.urlPath

	var payload Request
	payload.VoiceID = c.voice
	payload.Prompt = c.prompt
	payload.Text = c.text
	payload.Speed = config.Coqui.TTSSpeed

	b, _ := json.Marshal(&payload)

	req, _ := http.NewRequest("POST", u.String(), bytes.NewReader(b))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	return req
}

func (c TTSOptionsCoqui) WithSimilarityBoost(float32) common.TTSOptions {
	return c
}

func (c TTSOptionsCoqui) WithStability(float32) common.TTSOptions {
	return c
}

func (c TTSOptionsCoqui) WithText(text string) common.TTSOptions {
	c.text = text

	return c
}

func (c TTSOptionsCoqui) WithVoice(voice string) common.TTSOptions {
	c.voice = voice

	return c
}

func (c TTSOptionsCoqui) WithAudioEncoding(string) common.TTSOptions {
	return c
}

func (c TTSOptionsCoqui) WithSpeakingRate(float64) common.TTSOptions {
	return c
}

func (c TTSOptionsCoqui) WithPitch(float64) common.TTSOptions {
	return c
}

func (c TTSOptionsCoqui) WithVolumeGain(float64) common.TTSOptions {
	return c
}

func (c TTSOptionsCoqui) WithLanguage(lang string) common.TTSOptions {
	c.language = lang

	return c
}

func (c TTSOptionsCoqui) WithTimeout(timeout time.Duration) common.TTSOptions {
	c.timeout = timeout

	return c
}

func (c TTSOptionsCoqui) GetTimeout() time.Duration {
	return c.timeout
}

func (c TTSOptionsCoqui) WithModel(model common.TTSModel) common.TTSOptions {
	switch model {
	case CoquiV1TTS:
		c.urlPath = "/api/v2/samples"
	case XTTS:
		c.urlPath = "/api/v2/samples/xtts/render/"
	}

	return c
}

func (c TTSOptionsCoqui) WithPrompt(prompt string) common.TTSOptions {
	c.prompt = prompt

	return c
}
