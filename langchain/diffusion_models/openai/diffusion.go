package openai

import (
	"context"
	"encoding/base64"
	"github.com/kamushadenes/chloe/langchain/diffusion_models/common"
	"github.com/kamushadenes/chloe/logging"
	"github.com/sashabaranov/go-openai"
	"io"
	"net/http"
)

type DiffusionOpenAI struct {
	client *openai.Client
}

func NewDiffusionOpenAI(token string) *DiffusionOpenAI {
	return &DiffusionOpenAI{client: openai.NewClient(token)}
}

func (d *DiffusionOpenAI) Generate(message common.DiffusionMessage) (common.DiffusionResult, error) {
	return d.GenerateWithContext(context.Background(), message)
}

func (d *DiffusionOpenAI) GenerateWithContext(ctx context.Context, message common.DiffusionMessage) (common.DiffusionResult, error) {
	opts := NewDiffusionOptionsOpenAI().
		WithPrompt(message.Prompt)

	return d.GenerateWithOptions(ctx, opts)
}

func (d *DiffusionOpenAI) GenerateWithOptions(ctx context.Context, opts common.DiffusionOptions) (common.DiffusionResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := d.client.CreateImage(ctx, opts.GetRequest().(openai.ImageRequest))
	if err != nil {
		return common.DiffusionResult{}, err
	}

	var res common.DiffusionResult

	for k := range resp.Data {
		if resp.Data[k].URL != "" {
			hr, err := http.Get(resp.Data[k].URL)
			if err != nil {
				return common.DiffusionResult{}, err
			}
			defer hr.Body.Close()

			body, err := io.ReadAll(hr.Body)
			if err != nil {
				return common.DiffusionResult{}, err
			}

			res.Images = append(res.Images, body)
		} else if resp.Data[k].B64JSON != "" {
			data, err := base64.StdEncoding.DecodeString(resp.Data[k].B64JSON)
			if err != nil {
				return common.DiffusionResult{}, err
			}

			res.Images = append(res.Images, data)
		}
	}

	res.Usage = common.DiffusionUsage{
		ImageCount: opts.GetRequest().(openai.ImageRequest).N,
	}

	res.CalculateCosts(opts.GetModel())

	logger.Info().
		Str("provider", "openai").
		Str("model", opts.GetModel().Name).
		Float64("cost", res.Cost.TotalCost).
		Msg("image generation done")

	return res, nil
}
