package google

import (
	"time"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/common"
)

type TTSOptionsGoogle struct {
	req     *texttospeechpb.SynthesizeSpeechRequest
	timeout time.Duration
}

func NewTTSOptionsGoogle() common.TTSOptions {
	return &TTSOptionsGoogle{req: &texttospeechpb.SynthesizeSpeechRequest{}}
}

func (c TTSOptionsGoogle) GetRequest() interface{} {
	if c.req.Voice == nil {
		c.req.Voice = &texttospeechpb.VoiceSelectionParams{}
	}

	if c.req.Voice.LanguageCode == "" {
		c.req.Voice.LanguageCode = config.Google.TTSLanguageCode
	}

	if c.req.Voice.Name == "" {
		c.req.Voice.Name = config.Google.TTSVoiceName
	}

	if c.req.AudioConfig == nil {
		c.req.AudioConfig = &texttospeechpb.AudioConfig{}
	}

	if c.req.AudioConfig.AudioEncoding == 0 {
		c.req.AudioConfig.AudioEncoding = config.Google.TTSEncoding
	}

	if c.req.AudioConfig.SpeakingRate == 0 {
		c.req.AudioConfig.SpeakingRate = config.Google.TTSSpeakingRate
	}

	if c.req.AudioConfig.Pitch == 0 {
		c.req.AudioConfig.Pitch = config.Google.TTSPitch
	}

	if c.req.AudioConfig.VolumeGainDb == 0 {
		c.req.AudioConfig.VolumeGainDb = config.Google.TTSVolumeGain
	}

	return c.req
}

func (c TTSOptionsGoogle) WithSimilarityBoost(sb float32) common.TTSOptions {
	return c
}

func (c TTSOptionsGoogle) WithStability(st float32) common.TTSOptions {
	return c
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
