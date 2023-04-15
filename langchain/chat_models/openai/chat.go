package openai

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/chat_models"
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/kamushadenes/chloe/tokenizer"
	"github.com/sashabaranov/go-openai"
	"io"
)

type ChatOpenAI struct {
	client *openai.Client
	model  string
}

func NewChatOpenAI(token string, model string) chat_models.Chat {
	return &ChatOpenAI{client: openai.NewClient(token), model: model}
}

func NewChatOpenAIWithDefaultModel(token string) chat_models.Chat {
	return NewChatOpenAI(token, "gpt-3.5-turbo")
}

func (c *ChatOpenAI) Chat(messages ...schema.Message) (schema.ChatResult, error) {
	return c.ChatWithContext(context.Background(), messages...)
}

func (c *ChatOpenAI) ChatWithContext(ctx context.Context, messages ...schema.Message) (schema.ChatResult, error) {
	opts := NewChatOptionsOpenAI().WithMessages(messages).WithModel(c.model)

	return c.ChatWithOptions(ctx, opts)
}

func (c *ChatOpenAI) ChatWithOptions(ctx context.Context, opts schema.ChatOptions) (schema.ChatResult, error) {

	resp, err := c.client.CreateChatCompletion(ctx, opts.GetRequest())
	if err != nil {
		return schema.ChatResult{}, err
	}

	var res schema.ChatResult

	for k := range resp.Choices {
		res.Generations = append(res.Generations, schema.ChatGeneration{
			Text: resp.Choices[k].Message.Content,
			Message: schema.Message{
				Name:    resp.Choices[k].Message.Name,
				Role:    schema.Role(resp.Choices[k].Message.Role),
				Content: resp.Choices[k].Message.Content,
			},
		})
	}

	res.Usage = schema.ChatUsage{
		PromptTokens:     resp.Usage.PromptTokens,
		CompletionTokens: resp.Usage.CompletionTokens,
		TotalTokens:      resp.Usage.TotalTokens,
	}

	return res, nil
}

func (c *ChatOpenAI) ChatStream(w io.Writer, messages ...schema.Message) (schema.ChatResult, error) {
	return c.ChatStreamWithContext(context.Background(), w, messages...)
}

func (c *ChatOpenAI) ChatStreamWithContext(ctx context.Context, w io.Writer, messages ...schema.Message) (schema.ChatResult, error) {
	opts := NewChatOptionsOpenAI().WithMessages(messages).WithModel(c.model)

	return c.ChatStreamWithOptions(ctx, w, opts)
}

func (c *ChatOpenAI) ChatStreamWithOptions(ctx context.Context, w io.Writer, opts schema.ChatOptions) (schema.ChatResult, error) {

	stream, err := c.client.CreateChatCompletionStream(ctx, opts.GetRequest())
	if err != nil {
		return schema.ChatResult{}, err
	}
	defer stream.Close()

	var res schema.ChatResult
	res.Usage = schema.ChatUsage{}
	res.Generations[0] = schema.ChatGeneration{}

	msgs := opts.GetMessages()
	for k := range msgs {
		res.Usage.PromptTokens += tokenizer.CountTokens(c.model, msgs[k].Content)
	}

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			res.Generations[0].Message = schema.AssistantMessage(res.Generations[0].Text)

			return res, nil
		}

		if err != nil {
			return res, err
		}

		res.Generations[0].Text += resp.Choices[0].Delta.Content
		res.Usage.CompletionTokens += tokenizer.CountTokens(c.model, resp.Choices[0].Delta.Content)
		res.Usage.TotalTokens += tokenizer.CountTokens(c.model, resp.Choices[0].Delta.Content)

		if _, err := w.Write([]byte(resp.Choices[0].Delta.Content)); err != nil {
			return res, err
		}
	}
}
