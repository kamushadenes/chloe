package google

import (
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"time"
)

type TTSOptionsGoogle struct {
	req     *texttospeechpb.SynthesizeSpeechRequest
	timeout time.Duration
}

func NewTTSOptionsGoogle() common.TTSOptions {
	return &TTSOptionsGoogle{req: &texttospeechpb.SynthesizeSpeechRequest{}}
}

func (c TTSOptionsGoogle) GetRequest() interface{} {
	return c.req
}

func (c TTSOptionsGoogle) WithText(text string) common.TTSOptions {
	c.req.Input = &texttospeechpb.SynthesisInput{
		InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
	}
	return c
}

func (c TTSOptionsGoogle) WithVoice(voice string) common.TTSOptions {
	if c.req.Voice == nil {
		c.req.Voice = &texttospeechpb.VoiceSelectionParams{}
	}

	c.req.Voice.Name = voice

	return c
}

func (c TTSOptionsGoogle) WithAudioEncoding(encoding string) common.TTSOptions {
	if c.req.AudioConfig == nil {
		c.req.AudioConfig = &texttospeechpb.AudioConfig{}
	}

	c.req.AudioConfig.AudioEncoding = texttospeechpb.AudioEncoding(texttospeechpb.AudioEncoding_value[encoding])

	return c
}

func (c TTSOptionsGoogle) WithSpeakingRate(rate float64) common.TTSOptions {
	if c.req.AudioConfig == nil {
		c.req.AudioConfig = &texttospeechpb.AudioConfig{}
	}

	c.req.AudioConfig.SpeakingRate = rate

	return c
}

func (c TTSOptionsGoogle) WithPitch(pitch float64) common.TTSOptions {
	if c.req.AudioConfig == nil {
		c.req.AudioConfig = &texttospeechpb.AudioConfig{}
	}

	c.req.AudioConfig.Pitch = pitch

	return c
}

func (c TTSOptionsGoogle) WithVolumeGain(gain float64) common.TTSOptions {
	if c.req.AudioConfig == nil {
		c.req.AudioConfig = &texttospeechpb.AudioConfig{}
	}

	c.req.AudioConfig.VolumeGainDb = gain

	return c
}

func (c TTSOptionsGoogle) WithLanguage(language string) common.TTSOptions {
	if c.req.Voice == nil {
		c.req.Voice = &texttospeechpb.VoiceSelectionParams{}
	}

	c.req.Voice.LanguageCode = language

	return c
}

func (c TTSOptionsGoogle) WithTimeout(timeout time.Duration) common.TTSOptions {
	c.timeout = timeout
	return c
}

func (c TTSOptionsGoogle) GetTimeout() time.Duration {
	return c.timeout
}
