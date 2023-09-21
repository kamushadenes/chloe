package elevenlabs

import (
	"time"

	elevenlabs "github.com/haguro/elevenlabs-go"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/common"
)

type TTSOptionsElevenLabs struct {
	req     elevenlabs.TextToSpeechRequest
	timeout time.Duration
	voice   string
}

func NewTTSOptionsElevenLabs() common.TTSOptions {
	return &TTSOptionsElevenLabs{
		req:   elevenlabs.TextToSpeechRequest{},
		voice: config.ElevenLabs.TTSVoiceName,
	}
}

func (c TTSOptionsElevenLabs) GetRequest() interface{} {
	if c.req.ModelID == "" {
		c.req.ModelID = config.ElevenLabs.TTSLanguageCode
	}

	return c.req
}

func (c TTSOptionsElevenLabs) WithSimilarityBoost(sb float32) common.TTSOptions {
	c.req.VoiceSettings.SimilarityBoost = sb

	return c
}

func (c TTSOptionsElevenLabs) WithStability(st float32) common.TTSOptions {
	c.req.VoiceSettings.Stability = st

	return c
}

func (c TTSOptionsElevenLabs) WithText(text string) common.TTSOptions {
	c.req.Text = text

	return c
}

func (c TTSOptionsElevenLabs) WithVoice(voice string) common.TTSOptions {
	c.voice = voice

	return c
}

func (c TTSOptionsElevenLabs) WithAudioEncoding(string) common.TTSOptions {
	return c
}

func (c TTSOptionsElevenLabs) WithSpeakingRate(float64) common.TTSOptions {
	return c
}

func (c TTSOptionsElevenLabs) WithPitch(float64) common.TTSOptions {
	return c
}

func (c TTSOptionsElevenLabs) WithVolumeGain(float64) common.TTSOptions {
	return c
}

func (c TTSOptionsElevenLabs) WithLanguage(language string) common.TTSOptions {
	c.req.ModelID = language

	return c
}

func (c TTSOptionsElevenLabs) WithTimeout(timeout time.Duration) common.TTSOptions {
	elevenlabs.SetTimeout(timeout)

	c.timeout = timeout

	return c
}

func (c TTSOptionsElevenLabs) GetTimeout() time.Duration {
	return c.timeout
}

func (c TTSOptionsElevenLabs) WithModel(common.TTSModel) common.TTSOptions {
	return c
}

func (c TTSOptionsElevenLabs) WithPrompt(string) common.TTSOptions {
	return c
}
