package openai

import (
	"context"
	"io"

	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/writer_structs"
	"github.com/kamushadenes/chloe/tokenizer"
	"github.com/sashabaranov/go-openai"
)

func (c *ChatOpenAI) ChatStream(w io.Writer, messages ...messages.Message) (common.ChatResult, error) {
	return c.ChatStreamWithContext(context.Background(), w, messages...)
}

func (c *ChatOpenAI) ChatStreamWithContext(ctx context.Context, w io.Writer, messages ...messages.Message) (common.ChatResult, error) {
	opts := NewChatOptionsOpenAI().WithMessages(messages).WithModel(c.Model.Name)

	return c.ChatStreamWithOptions(ctx, w, opts)
}

func (c *ChatOpenAI) ChatStreamWithOptions(ctx context.Context, w io.Writer, opts common.ChatOptions) (common.ChatResult, error) {
	logger := logging.GetLogger()

	msgs, err := c.LoadUserMessages(ctx)
	if err != nil {
		return common.ChatResult{}, err
	}

	msgs = append(msgs, opts.GetMessages()...)
	msgs = c.ReduceTokens(opts.GetSystemMessages(), msgs)

	opts = opts.WithMessages(msgs)

	stream, err := c.Client.CreateChatCompletionStream(ctx, opts.GetRequest().(openai.ChatCompletionRequest))
	if err != nil {
		return common.ChatResult{}, err
	}
	defer stream.Close()

	var res common.ChatResult
	res.Usage = common.ChatUsage{}

	modelName := c.Model.Tokenizer
	if c.Model.Tokenizer == "" {
		modelName = c.Model.Name
	}

	for k := range msgs {
		res.Usage.PromptTokens += tokenizer.CountTokens(modelName, msgs[k].Content)
	}

	var fnc functions.FunctionCall

	for {
		resp, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			for k := range res.Generations {
				m := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "internal")
				m.Context = ctx
				m.Role = string(messages.Assistant)
				m.User = c.User
				m.SetContent(res.Generations[k].Text)
				if err := m.Save(ctx); err != nil {
					return common.ChatResult{}, err
				}

				res.Generations[k].Message = messages.AssistantMessage(res.Generations[k].Text)
				res.Generations[k].Message.FunctionCall = &fnc
			}

			res.CalculateCosts(c.Model)

			logger.Info().
				Str("provider", "openai").
				Str("model", c.Model.Name).
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

			if resp.Choices[k].Delta.FunctionCall != nil {
				if resp.Choices[k].Delta.FunctionCall.Name != "" {
					fnc.Name = resp.Choices[k].Delta.FunctionCall.Name
				}
				if resp.Choices[k].Delta.FunctionCall.Arguments != "" {
					fnc.Arguments += resp.Choices[k].Delta.FunctionCall.Arguments
				}
			}
			res.Generations[k].Text += resp.Choices[k].Delta.Content
			res.Usage.CompletionTokens += tokenizer.CountTokens(modelName, resp.Choices[k].Delta.Content)
			res.Usage.TotalTokens += tokenizer.CountTokens(modelName, resp.Choices[k].Delta.Content)

			if resp.Choices[k].FinishReason != "" {
				res.Generations[k].FinishReason = string(resp.Choices[k].FinishReason)
			}

			if _, err := w.Write([]byte(resp.Choices[k].Delta.Content)); err != nil {
				return res, err
			} else {
				w.(writer_structs.ChloeWriter).Flush()
			}
		}
	}
}
