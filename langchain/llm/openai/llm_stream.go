package openai

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/llm"
	"github.com/kamushadenes/chloe/tokenizer"
	"io"
)

func (c *LLMOpenAI) GenerateStream(w io.Writer, prompt ...string) (llm.LLMResult, error) {
	return c.GenerateStreamWithContext(context.Background(), w, prompt...)
}

func (c *LLMOpenAI) GenerateStreamWithContext(ctx context.Context, w io.Writer, prompt ...string) (llm.LLMResult, error) {
	opts := NewLLMOptionsOpenAI().WithPrompt(prompt).WithModel(c.model.Name)

	return c.GenerateStreamWithOptions(ctx, w, opts)
}

func (c *LLMOpenAI) GenerateStreamWithOptions(ctx context.Context, w io.Writer, opts llm.LLMOptions) (llm.LLMResult, error) {
	stream, err := c.client.CreateCompletionStream(ctx, opts.GetRequest())
	if err != nil {
		return llm.LLMResult{}, err
	}
	defer stream.Close()

	var res llm.LLMResult
	res.Usage = llm.LLMUsage{}
	res.Generations[0] = llm.LLMGeneration{}

	msgs := opts.GetPrompt()
	for k := range msgs {
		res.Usage.PromptTokens += tokenizer.CountTokens(c.model.Name, msgs[k])
	}

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			for k := range res.Generations {
				res.Generations[k].FinishReason = resp.Choices[k].FinishReason
			}

			return res, nil
		}

		if err != nil {
			return res, err
		}

		for k := range resp.Choices {
			if len(res.Generations) <= k {
				res.Generations = append(res.Generations, llm.LLMGeneration{})
			}

			res.Generations[k].Index = resp.Choices[k].Index
			res.Generations[k].Text += resp.Choices[k].Text

			res.Usage.CompletionTokens += tokenizer.CountTokens(c.model.Name, resp.Choices[k].Text)
			res.Usage.TotalTokens += tokenizer.CountTokens(c.model.Name, resp.Choices[k].Text)

			if _, err := w.Write([]byte(resp.Choices[k].Text)); err != nil {
				return res, err
			}
		}
	}
}
