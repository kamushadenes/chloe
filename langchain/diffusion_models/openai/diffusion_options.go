package openai

import (
	"github.com/kamushadenes/chloe/langchain/diffusion_models/common"
	"github.com/sashabaranov/go-openai"
	"time"
)

type DiffusionOptionsOpenAI struct {
	req     openai.ImageRequest
	model   *common.DiffusionModel
	timeout time.Duration
}

func NewDiffusionOptionsOpenAI() common.DiffusionOptions {
	opts := DiffusionOptionsOpenAI{req: openai.ImageRequest{}}

	return opts.WithCount(1).WithModel(DallE512X512)
}

func (c DiffusionOptionsOpenAI) GetRequest() interface{} {
	return c.req
}

func (c DiffusionOptionsOpenAI) WithPrompt(prompt string) common.DiffusionOptions {
	c.req.Prompt = prompt
	return c
}

func (c DiffusionOptionsOpenAI) WithModel(model *common.DiffusionModel) common.DiffusionOptions {
	c.model = model
	return c
}

func (c DiffusionOptionsOpenAI) GetModel() *common.DiffusionModel {
	return c.model
}

func (c DiffusionOptionsOpenAI) WithTimeout(timeout time.Duration) common.DiffusionOptions {
	c.timeout = timeout
	return c
}

func (c DiffusionOptionsOpenAI) GetTimeout() time.Duration {
	return c.timeout
}

func (c DiffusionOptionsOpenAI) WithCount(count int) common.DiffusionOptions {
	c.req.N = count
	return c
}

func (c DiffusionOptionsOpenAI) WithFormat(format string) common.DiffusionOptions {
	c.req.ResponseFormat = format
	return c
}
