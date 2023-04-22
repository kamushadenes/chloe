package google

import (
	"bytes"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/logging"
	"io"
)

type TTSGoogle struct{}

func NewTTSGoogle() *TTSGoogle {
	return &TTSGoogle{}
}

func (c *TTSGoogle) TTS(message common.TTSMessage) (common.TTSResult, error) {
	return c.TTSWithContext(context.Background(), message)
}

func (c *TTSGoogle) TTSWithContext(ctx context.Context, message common.TTSMessage) (common.TTSResult, error) {
	opts := NewTTSOptionsGoogle().
		WithText(message.Text).
		WithTimeout(config.Timeouts.TTS)

	return c.TTSWithOptions(ctx, opts)
}

func (c *TTSGoogle) TTSWithOptions(ctx context.Context, opts common.TTSOptions) (common.TTSResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		return common.TTSResult{}, err
	}

	defer func(client *texttospeech.Client) {
		_ = client.Close()
	}(client)

	resp, err := client.SynthesizeSpeech(ctx, opts.GetRequest().(*texttospeechpb.SynthesizeSpeechRequest))
	if err != nil {
		return common.TTSResult{}, errors.Wrap(errors.ErrTTSFailed, err)
	}

	var res common.TTSResult

	w := bytes.NewBuffer(nil)

	if _, err := io.Copy(w, bytes.NewReader(resp.AudioContent)); err != nil {
		return common.TTSResult{}, errors.Wrap(errors.ErrTTSFailed, err)
	}

	res.Audio = w.Bytes()

	res.ContentType = "application/octet-stream"

	switch opts.GetRequest().(*texttospeechpb.SynthesizeSpeechRequest).AudioConfig.AudioEncoding {
	case texttospeechpb.AudioEncoding_MP3:
		res.ContentType = "audio/mpeg"
	case texttospeechpb.AudioEncoding_OGG_OPUS:
		res.ContentType = "audio/ogg"
	case texttospeechpb.AudioEncoding_LINEAR16:
		res.ContentType = "audio/wav"
	case texttospeechpb.AudioEncoding_MULAW:
		res.ContentType = "audio/mulaw"
	case texttospeechpb.AudioEncoding_ALAW:
		res.ContentType = "audio/alaw"
	}

	logger.Info().
		Str("provider", "google").
		Msg("tts done")

	return res, nil
}
