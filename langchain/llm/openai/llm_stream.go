package openai

import (
	"context"
	"io"

	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/llm/common"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/tokenizer"
	openai "github.com/sashabaranov/go-openai"
)

func (c *LLMOpenAI) GenerateStream(w io.Writer, prompt ...string) (common.LLMResult, error) {
	return c.GenerateStreamWithContext(context.Background(), w, prompt...)
}

func (c *LLMOpenAI) GenerateStreamWithContext(ctx context.Context, w io.Writer, prompt ...string) (common.LLMResult, error) {
	opts := NewLLMOptionsOpenAI().WithPrompt(prompt).WithModel(c.Model.Name)

	return c.GenerateStreamWithOptions(ctx, w, opts)
}

func (c *LLMOpenAI) GenerateStreamWithOptions(ctx context.Context, w io.Writer, opts common.LLMOptions) (common.LLMResult, error) {
	logger := logging.GetLogger()

	stream, err := c.Client.CreateCompletionStream(ctx, opts.GetRequest().(openai.CompletionRequest))
	if err != nil {
		return common.LLMResult{}, err
	}
	defer stream.Close()

	var res common.LLMResult
	res.Usage = common.LLMUsage{}
	res.Generations[0] = common.LLMGeneration{}

	modelName := c.Model.Tokenizer
	if c.Model.Tokenizer == "" {
		modelName = c.Model.Name
	}

	msgs := opts.GetPrompt()
	for k := range msgs {
		res.Usage.PromptTokens += tokenizer.CountTokens(modelName, msgs[k])
	}

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			for k := range res.Generations {
				res.Generations[k].FinishReason = resp.Choices[k].FinishReason
			}

			res.CalculateCosts(c.Model)

			logger.Info().
				Str("provider", "openai").
				Str("model", c.Model.Name).
				Float64("cost", res.Cost.TotalCost).
				Int("tokens", res.Usage.TotalTokens).
				Msg("llm stream done")

			return res, nil
		}

		if err != nil {
			return res, err
		}

		for k := range resp.Choices {
			if len(res.Generations) <= k {
				res.Generations = append(res.Generations, common.LLMGeneration{})
			}

			res.Generations[k].Index = resp.Choices[k].Index
			res.Generations[k].Text += resp.Choices[k].Text

			res.Usage.CompletionTokens += tokenizer.CountTokens(modelName, resp.Choices[k].Text)
			res.Usage.TotalTokens += tokenizer.CountTokens(modelName, resp.Choices[k].Text)

			if _, err := w.Write([]byte(resp.Choices[k].Text)); err != nil {
				return res, err
			}
		}
	}
}
