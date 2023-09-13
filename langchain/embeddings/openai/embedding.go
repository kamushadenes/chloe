package openai

import (
	"context"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/embeddings/common"
	"github.com/kamushadenes/chloe/logging"
	"github.com/sashabaranov/go-openai"
)

type EmbeddingOpenAI struct {
	Client *openai.Client
	Model  *common.EmbeddingModel
}

func NewEmbeddingOpenAI(token string, model *common.EmbeddingModel) common.Embedding {
	return &EmbeddingOpenAI{Client: openai.NewClient(token), Model: model}
}

func NewEmbeddingOpenAIWithDefaultModel(token string) common.Embedding {
	return NewEmbeddingOpenAI(token, AdaEmbeddingV2)
}

func (c *EmbeddingOpenAI) Embed(text []string) (common.EmbeddingResult, error) {
	return c.EmbedWithContext(context.Background(), text)
}

func (c *EmbeddingOpenAI) EmbedWithContext(ctx context.Context, text []string) (common.EmbeddingResult, error) {

	opts := NewEmbeddingOptionsOpenAI().
		WithText(text).
		WithModel(c.Model.Name).
		WithTimeout(config.Timeouts.Completion)

	return c.EmbedWithOptions(ctx, opts)
}

func (c *EmbeddingOpenAI) EmbedWithOptions(ctx context.Context, opts common.EmbeddingOptions) (common.EmbeddingResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	resp, err := c.Client.CreateEmbeddings(ctx, opts.GetRequest().(openai.EmbeddingRequest))
	if err != nil {
		logger.Error().
			Str("provider", "openai").
			Str("model", c.Model.Name).
			Err(err).
			Msg("embedding error")
		return common.EmbeddingResult{}, err
	}

	var res common.EmbeddingResult

	for k := range resp.Data {
		var e common.EmbeddingObject
		e.Object = resp.Data[k].Object
		e.Embedding = resp.Data[k].Embedding
		e.Index = resp.Data[k].Index

		res.Embeddings = append(res.Embeddings, e)
	}

	res.Usage = common.EmbeddingUsage{
		PromptTokens: resp.Usage.PromptTokens,
		TotalTokens:  resp.Usage.TotalTokens,
	}

	res.CalculateCosts(c.Model)

	logger.Info().
		Str("provider", "openai").
		Str("model", c.Model.Name).
		Float64("cost", res.Cost.TotalCost).
		Int("tokens", res.Usage.TotalTokens).
		Msg("embedding done")

	return res, nil
}
