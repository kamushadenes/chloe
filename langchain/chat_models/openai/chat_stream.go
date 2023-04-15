package openai

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/kamushadenes/chloe/tokenizer"
	"io"
)

func (c *ChatOpenAI) ChatStream(w io.Writer, messages ...chat_models.Message) (chat_models.ChatResult, error) {
	return c.ChatStreamWithContext(context.Background(), w, messages...)
}

func (c *ChatOpenAI) ChatStreamWithContext(ctx context.Context, w io.Writer, messages ...chat_models.Message) (chat_models.ChatResult, error) {
	opts := NewChatOptionsOpenAI().WithMessages(messages).WithModel(c.model.Name)

	return c.ChatStreamWithOptions(ctx, w, opts)
}

func (c *ChatOpenAI) ChatStreamWithOptions(ctx context.Context, w io.Writer, opts chat_models.ChatOptions) (chat_models.ChatResult, error) {
	stream, err := c.client.CreateChatCompletionStream(ctx, opts.GetRequest())
	if err != nil {
		return chat_models.ChatResult{}, err
	}
	defer stream.Close()

	var res chat_models.ChatResult
	res.Usage = chat_models.ChatUsage{}
	res.Generations[0] = chat_models.ChatGeneration{}

	msgs := opts.GetMessages()
	for k := range msgs {
		res.Usage.PromptTokens += tokenizer.CountTokens(c.model.Name, msgs[k].Content)
	}

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			for k := range res.Generations {
				res.Generations[k].Message = chat_models.AssistantMessage(res.Generations[k].Text)
			}

			res.CalculateCosts(c.model)

			return res, nil
		}

		if err != nil {
			return res, err
		}

		for k := range resp.Choices {
			if len(res.Generations) <= k {
				res.Generations = append(res.Generations, chat_models.ChatGeneration{})
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
