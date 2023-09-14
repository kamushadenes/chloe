package whispercpp

import (
	"context"
	"fmt"
	"github.com/ggerganov/whisper.cpp/bindings/go/pkg/whisper"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/audio_models/common"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/media"
	"io"
)

type ASRWhisperCpp struct {
	model *common.ASRModel
}

func NewASRWhisperCpp(model *common.ASRModel) common.ASR {
	return &ASRWhisperCpp{
		model: model,
	}
}

func NewASRWhisperCppWithDefaultModel() common.ASR {
	return NewASRWhisperCpp(Medium)
}

func (c *ASRWhisperCpp) Transcribe(audioFile string) (common.ASRResult, error) {
	return c.TranscribeWithContext(context.Background(), audioFile)
}

func (c *ASRWhisperCpp) TranscribeWithContext(ctx context.Context, audioFile string) (common.ASRResult, error) {
	opts := NewASROptionsWhisperCpp().
		WithAudioFile(audioFile).
		WithLanguage(config.WhisperCpp.Language).
		WithModel(c.model.Name)

	return c.TranscribeWithOptions(ctx, opts)
}

func (c *ASRWhisperCpp) TranscribeWithOptions(ctx context.Context, opts common.ASROptions) (common.ASRResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	wctx, err := opts.GetRequest().(whisper.Model).NewContext()
	if err != nil {
		return common.ASRResult{}, err
	}

	wctx.SetTranslate(false)
	if err := wctx.SetLanguage(opts.(ASROptionsWhisperCpp).language); err != nil {
		return common.ASRResult{}, err
	}

	var cb whisper.SegmentCallback
	wctx.ResetTimings()
	if err := wctx.Process(opts.(ASROptionsWhisperCpp).audioSamples, cb, nil); err != nil {
		return common.ASRResult{}, err
	}

	var res common.ASRResult
	switch opts.(ASROptionsWhisperCpp).outFormat {
	case "srt":
		n := 1
		for {
			segment, err := wctx.NextSegment()
			if err == io.EOF {
				break
			} else if err != nil {
				break
			}
			res.Text += fmt.Sprintf("%d\n", n)
			res.Text += fmt.Sprintf("%s --> %s\n", srtTimestamp(segment.Start), srtTimestamp(segment.End))
			res.Text += fmt.Sprintf("%s\n", segment.Text)
			res.Text += "\n"
			n++
		}
	default:
		for {
			segment, err := wctx.NextSegment()
			if err == io.EOF {
				break
			} else if err != nil {
				break
			}
			res.Text += fmt.Sprintf(" %s", segment.Text)
		}
	}

	dur, err := media.GetMediaDuration(opts.GetAudioFile())
	if err != nil {
		logger.Error().Err(err).Msg("error getting media duration")
	}

	res.Usage = common.ASRUsage{
		Duration: dur,
	}

	res.CalculateCosts(c.model)

	logger.Info().
		Str("provider", "whispercpp").
		Str("model", c.model.Name).
		Float64("cost", res.Cost.TotalCost).
		Dur("duration", res.Usage.Duration).
		Msg("audio transcription done")

	return res, nil
}

func (c *ASRWhisperCpp) Translate(audioFile string) (common.ASRResult, error) {
	return c.TranslateWithContext(context.Background(), audioFile)
}

func (c *ASRWhisperCpp) TranslateWithContext(ctx context.Context, audioFile string) (common.ASRResult, error) {
	opts := NewASROptionsWhisperCpp().
		WithAudioFile(audioFile).
		WithModel(c.model.Name).
		WithTimeout(config.Timeouts.Transcription)

	return c.TranslateWithOptions(ctx, opts)
}

func (c *ASRWhisperCpp) TranslateWithOptions(ctx context.Context, opts common.ASROptions) (common.ASRResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	wctx, err := opts.(ASROptionsWhisperCpp).loadedModel.NewContext()
	if err != nil {
		return common.ASRResult{}, err
	}

	wctx.SetTranslate(true)
	if err := wctx.SetLanguage(opts.(ASROptionsWhisperCpp).language); err != nil {
		return common.ASRResult{}, err
	}

	var cb whisper.SegmentCallback
	wctx.ResetTimings()
	if err := wctx.Process(opts.(ASROptionsWhisperCpp).audioSamples, cb, nil); err != nil {
		return common.ASRResult{}, err
	}

	var res common.ASRResult
	switch opts.(ASROptionsWhisperCpp).outFormat {
	case "srt":
		n := 1
		for {
			segment, err := wctx.NextSegment()
			if err == io.EOF {
				break
			} else if err != nil {
				break
			}
			res.Text += fmt.Sprintf("%d\n", n)
			res.Text += fmt.Sprintf("%s --> %s\n", srtTimestamp(segment.Start), srtTimestamp(segment.End))
			res.Text += fmt.Sprintf("%s\n", segment.Text)
			res.Text += "\n"
			n++
		}
	default:
		for {
			segment, err := wctx.NextSegment()
			if err == io.EOF {
				break
			} else if err != nil {
				break
			}
			res.Text += fmt.Sprintf(" %s", segment.Text)
		}
	}

	dur, err := media.GetMediaDuration(opts.GetAudioFile())
	if err != nil {
		return common.ASRResult{}, err
	}

	res.Usage = common.ASRUsage{
		Duration: dur,
	}

	res.CalculateCosts(c.model)

	logger.Info().
		Str("provider", "whispercpp").
		Str("model", c.model.Name).
		Float64("cost", res.Cost.TotalCost).
		Dur("duration", res.Usage.Duration).
		Msg("audio translation done")

	return res, nil
}
