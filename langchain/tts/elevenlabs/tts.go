package elevenlabs

import (
	"context"
	elevenlabs "github.com/haguro/elevenlabs-go"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/logging"
)

type TTSElevenLabs struct{}

func NewTTSElevenLabs() *TTSElevenLabs {
	return &TTSElevenLabs{}
}

func (c *TTSElevenLabs) TTS(message common.TTSMessage) (common.TTSResult, error) {
	return c.TTSWithContext(context.Background(), message)
}

func (c *TTSElevenLabs) TTSWithContext(ctx context.Context, message common.TTSMessage) (common.TTSResult, error) {
	opts := NewTTSOptionsElevenLabs().
		WithText(message.Text).
		WithTimeout(config.Timeouts.TTS)

	return c.TTSWithOptions(ctx, opts)
}

func (c *TTSElevenLabs) TTSWithOptions(ctx context.Context, opts common.TTSOptions) (common.TTSResult, error) {
	logger := logging.GetLogger()

	elevenlabs.SetAPIKey(config.ElevenLabs.APIKey)

	audio, err := elevenlabs.TextToSpeech(opts.(TTSOptionsElevenLabs).voice, opts.GetRequest().(elevenlabs.TextToSpeechRequest))
	if err != nil {
		return common.TTSResult{}, err
	}

	var res common.TTSResult

	res.Audio = audio

	res.ContentType = "audio/mpeg"

	logger.Info().
		Str("provider", "elevenlabs").
		Msg("tts done")

	return res, nil
}
