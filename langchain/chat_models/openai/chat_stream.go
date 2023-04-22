package openai

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/tokenizer"
	"io"
)

func (c *ChatOpenAI) ChatStream(w io.Writer, messages ...common.Message) (common.ChatResult, error) {
	return c.ChatStreamWithContext(context.Background(), w, messages...)
}

func (c *ChatOpenAI) ChatStreamWithContext(ctx context.Context, w io.Writer, messages ...common.Message) (common.ChatResult, error) {
	opts := NewChatOptionsOpenAI().WithMessages(messages).WithModel(c.model.Name)

	return c.ChatStreamWithOptions(ctx, w, opts)
}

func (c *ChatOpenAI) ChatStreamWithOptions(ctx context.Context, w io.Writer, opts common.ChatOptions) (common.ChatResult, error) {
	logger := logging.GetLogger()

	stream, err := c.client.CreateChatCompletionStream(ctx, opts.GetRequest())
	if err != nil {
		return common.ChatResult{}, err
	}
	defer stream.Close()

	var res common.ChatResult
	res.Usage = common.ChatUsage{}
	res.Generations[0] = common.ChatGeneration{}

	msgs := opts.GetMessages()
	for k := range msgs {
		res.Usage.PromptTokens += tokenizer.CountTokens(c.model.Name, msgs[k].Content)
	}

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			for k := range res.Generations {
				res.Generations[k].Message = common.AssistantMessage(res.Generations[k].Text)
			}

			res.CalculateCosts(c.model)

			logger.Info().
				Str("provider", "openai").
				Str("model", c.model.Name).
				Float64("cost", res.Cost.TotalCost).
				Int("tokens", res.Usage.TotalTokens).
				Msg("chat stream done")

			return res, nil
		}

		if err != nil {
			return res, err
		}

		for k := range resp.Choices {
			if len(res.Generations) <= k {
				res.Generations = append(res.Generations, common.ChatGeneration{})
			}

			res.Generations[k].Text += resp.Choices[k].Delta.Content
			res.Usage.CompletionTokens += tokenizer.CountTokens(c.model.Name, resp.Choices[k].Delta.Content)
			res.Usage.TotalTokens += tokenizer.CountTokens(c.model.Name, resp.Choices[k].Delta.Content)

			if _, err := w.Write([]byte(resp.Choices[k].Delta.Content)); err != nil {
				return res, err
			}
		}
	}
}
