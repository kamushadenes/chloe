package openai

import (
	"time"

	"github.com/kamushadenes/chloe/langchain/embeddings/common"
	openai "github.com/sashabaranov/go-openai"
)

type EmbeddingOptionsOpenAI struct {
	req     openai.EmbeddingRequest
	timeout time.Duration
}

func NewEmbeddingOptionsOpenAI() common.EmbeddingOptions {
	opts := &EmbeddingOptionsOpenAI{req: openai.EmbeddingRequest{}}

	return opts
}

func (c EmbeddingOptionsOpenAI) GetRequest() interface{} {
	return c.req
}

func (c EmbeddingOptionsOpenAI) WithTimeout(dur time.Duration) common.EmbeddingOptions {
	c.timeout = dur
	return c
}

func (c EmbeddingOptionsOpenAI) GetTimeout() time.Duration {
	return c.timeout
}

func (c EmbeddingOptionsOpenAI) WithText(text []string) common.EmbeddingOptions {
	c.req.Input = text
	return c
}

func (c EmbeddingOptionsOpenAI) GetText() []string {
	return c.req.Input.([]string)
}

func (c EmbeddingOptionsOpenAI) WithChunkSize(size int) common.EmbeddingOptions {
	// c.req.ChunkSize = size
	return c
}

func (c EmbeddingOptionsOpenAI) WithMaxRetries(retries int) common.EmbeddingOptions {
	// c.req.MaxRetries = retries
	return c
}

func (c EmbeddingOptionsOpenAI) WithModel(model string) common.EmbeddingOptions {
	if m, ok := modelStringToEnum[model]; ok {
		c.req.Model = m
	}

	return c
}
