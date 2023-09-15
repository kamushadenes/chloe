package coqui

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/logging"
)

type TTSCoqui struct {
	apiKey string
}

func NewTTSCoqui(apiKey string) *TTSCoqui {
	return &TTSCoqui{
		apiKey: apiKey,
	}
}

func (c *TTSCoqui) TTS(message common.TTSMessage) (common.TTSResult, error) {
	return c.TTSWithContext(context.Background(), message)
}

func (c *TTSCoqui) TTSWithContext(ctx context.Context, message common.TTSMessage) (common.TTSResult, error) {
	opts := NewTTSOptionsCoqui().
		WithText(message.Text).
		WithTimeout(config.Timeouts.TTS)

	switch config.Coqui.TTSModel {
	case "coqui_v1":
		opts = opts.WithModel("coqui_v1")
	case "xtts":
		opts = opts.WithModel("xtts")
	}

	return c.TTSWithOptions(ctx, opts)
}

func (c *TTSCoqui) TTSWithOptions(ctx context.Context, opts common.TTSOptions) (common.TTSResult, error) {
	logger := logging.GetLogger()

	req := opts.GetRequest().(*http.Request).WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return common.TTSResult{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return common.TTSResult{}, err
	}

	var r Response
	if err := json.Unmarshal(body, &r); err != nil {
		return common.TTSResult{}, err
	}

	if _, err := url.Parse(r.AudioURL); err != nil {
		return common.TTSResult{}, err
	}

	resp, err = http.Get(r.AudioURL)
	if err != nil {
		return common.TTSResult{}, err
	}
	defer resp.Body.Close()

	audio, err := io.ReadAll(resp.Body)
	if err != nil {
		return common.TTSResult{}, err
	}

	var res common.TTSResult

	res.Audio = audio

	res.ContentType = "audio/wav"

	logger.Info().
		Str("provider", "coqui").
		Msg("tts done")

	return res, nil
}
