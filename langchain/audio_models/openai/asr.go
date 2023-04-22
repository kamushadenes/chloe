package openai

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/audio_models/common"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/media"
	"github.com/sashabaranov/go-openai"
)

type ASROpenAI struct {
	client *openai.Client
	model  *common.ASRModel
}

func NewASROpenAI(token string, model *common.ASRModel) *ASROpenAI {
	return &ASROpenAI{client: openai.NewClient(token), model: model}
}

func NewASROpenAIWithDefaultModel(token string) *ASROpenAI {
	return NewASROpenAI(token, Whisper1)
}

func (c *ASROpenAI) Transcribe(audioFile string) (common.ASRResult, error) {
	return c.TranscribeWithContext(context.Background(), audioFile)
}

func (c *ASROpenAI) TranscribeWithContext(ctx context.Context, audioFile string) (common.ASRResult, error) {
	opts := NewASROptionsOpenAI().
		WithAudioFile(audioFile).
		WithModel(c.model.Name)

	return c.TranscribeWithOptions(ctx, opts)
}

func (c *ASROpenAI) TranscribeWithOptions(ctx context.Context, opts common.ASROptions) (common.ASRResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := c.client.CreateTranscription(ctx, opts.GetRequest().(openai.AudioRequest))
	if err != nil {
		return common.ASRResult{}, err
	}

	var res common.ASRResult

	res.Text = resp.Text

	dur, err := media.GetMediaDuration(opts.GetAudioFile())
	if err != nil {
		return common.ASRResult{}, err
	}

	res.Usage = common.ASRUsage{
		Duration: dur,
	}

	res.CalculateCosts(c.model)

	logger.Info().
		Str("provider", "openai").
		Str("model", c.model.Name).
		Float64("cost", res.Cost.TotalCost).
		Dur("duration", res.Usage.Duration).
		Msg("audio transcription done")

	return res, nil
}

func (c *ASROpenAI) Translate(audioFile string) (common.ASRResult, error) {
	return c.TranslateWithContext(context.Background(), audioFile)
}

func (c *ASROpenAI) TranslateWithContext(ctx context.Context, audioFile string) (common.ASRResult, error) {
	opts := NewASROptionsOpenAI().
		WithAudioFile(audioFile).
		WithModel(c.model.Name).
		WithTimeout(config.Timeouts.Transcription)

	return c.TranslateWithOptions(ctx, opts)
}

func (c *ASROpenAI) TranslateWithOptions(ctx context.Context, opts common.ASROptions) (common.ASRResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := c.client.CreateTranslation(ctx, opts.GetRequest().(openai.AudioRequest))
	if err != nil {
		return common.ASRResult{}, err
	}

	var res common.ASRResult

	res.Text = resp.Text

	dur, err := media.GetMediaDuration(opts.GetAudioFile())
	if err != nil {
		return common.ASRResult{}, err
	}

	res.Usage = common.ASRUsage{
		Duration: dur,
	}

	res.CalculateCosts(c.model)

	logger.Info().
		Str("provider", "openai").
		Str("model", c.model.Name).
		Float64("cost", res.Cost.TotalCost).
		Dur("duration", res.Usage.Duration).
		Msg("audio translation done")

	return res, nil
}
