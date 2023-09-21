package openai

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/chat_models/common"
	"github.com/kamushadenes/chloe/langchain/chat_models/messages"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/writer_structs"
	openai "github.com/sashabaranov/go-openai"
)

type ChatOpenAI struct {
	Client *openai.Client
	Model  *common.ChatModel
	User   *memory.User
}

func NewChatOpenAI(token string, model *common.ChatModel, user *memory.User) common.Chat {
	return &ChatOpenAI{Client: openai.NewClient(token), Model: model, User: user}
}

func NewChatOpenAIWithDefaultModel(token string, user *memory.User) common.Chat {
	return NewChatOpenAI(token, GPT35Turbo, user)
}

func (c *ChatOpenAI) Chat(userMsg *memory.Message, messages ...messages.Message) (common.ChatResult, error) {
	return c.ChatWithContext(context.Background(), userMsg, messages...)
}

func (c *ChatOpenAI) ChatWithContext(ctx context.Context, userMsg *memory.Message, messages ...messages.Message) (common.ChatResult, error) {
	logger := logging.GetLogger()

	msgs, err := c.LoadUserMessages(ctx)
	if err != nil {
		logger.Error().
			Str("provider", "openai").
			Str("model", c.Model.Name).
			Err(err).
			Msg("chat completion error")

		return common.ChatResult{}, err
	}

	for k := range msgs {
		msg := msgs[k]

		found := false
		for kk := range messages {
			if msg.ID == messages[kk].ID {
				fmt.Println("AAA")
				fmt.Println(msg.ID)
				found = true
				break
			}
		}

		if !found {
			messages = append(messages, msg)
		}
	}

	opts := NewChatOptionsOpenAI().
		WithMessages(msgs).
		WithModel(c.Model.Name).
		WithTimeout(config.Timeouts.Completion).
		WithUserMessage(userMsg)

	return c.ChatWithOptions(ctx, opts)
}

func (c *ChatOpenAI) ChatWithOptions(ctx context.Context, opts common.ChatOptions) (common.ChatResult, error) {
	logger := logging.GetLogger()

	if opts.GetTimeout() > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, opts.GetTimeout())
		defer cancel()
	}

	if opts.GetFunctions() == nil {
		acts := actions.GetAllowedActions()

		var defs []*functions.FunctionDefinition

		for k := range acts {
			defs = append(defs, acts[k].GetSchema())
		}

		opts = opts.WithFunctions(defs)
	}

	var fmsgs []messages.Message
	omsgs := opts.GetMessages()
	for k := range omsgs {
		if omsgs[k].Role != messages.System {
			fmsgs = append(fmsgs, omsgs[k])
		}
	}

	msgs := c.ReduceTokens(opts.GetSystemMessages(), fmsgs)

	opts = opts.WithMessages(msgs)

	resp, err := c.Client.CreateChatCompletion(ctx, opts.GetRequest().(openai.ChatCompletionRequest))
	if err != nil {
		logger.Error().
			Str("provider", "openai").
			Str("model", c.Model.Name).
			Err(err).
			Msg("chat completion error")
		return common.ChatResult{}, err
	}

	var res common.ChatResult

	for k := range resp.Choices {
		m := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "internal")
		m.Context = ctx
		m.Role = string(messages.Assistant)
		m.User = c.User
		m.SetContent(resp.Choices[k].Message.Content)

		if resp.Choices[k].Message.FunctionCall != nil {
			m.FunctionCallName = resp.Choices[k].Message.FunctionCall.Name
			m.FunctionCallArguments = resp.Choices[k].Message.FunctionCall.Arguments
		}

		if err := m.Save(ctx); err != nil {
			return common.ChatResult{}, err
		}

		g := common.ChatGeneration{
			FinishReason: string(resp.Choices[k].FinishReason),
			Text:         resp.Choices[k].Message.Content,
			Message: messages.Message{
				ID:      m.ID,
				Name:    resp.Choices[k].Message.Name,
				Role:    messages.Role(resp.Choices[k].Message.Role),
				Content: resp.Choices[k].Message.Content,
			},
		}

		if resp.Choices[k].Message.FunctionCall != nil {
			g.Message.FunctionCall = &functions.FunctionCall{
				Name:      resp.Choices[k].Message.FunctionCall.Name,
				Arguments: resp.Choices[k].Message.FunctionCall.Arguments,
			}
		}

		res.Generations = append(res.Generations, g)
	}

	for k := range res.Generations {
		g := res.Generations[k]

		msgs = append(msgs, g.Message)

		if g.Message.FunctionCall != nil {
			w := writer_structs.NewMockWriter()

			req := action_structs.NewActionRequest()

			req.Context = ctx
			req.Action = g.Message.FunctionCall.Name
			req.Writer = w
			req.Message = opts.GetUserMessage()

			fm := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "internal")
			fm.Context = ctx
			fm.Role = string(messages.Assistant)
			fm.User = c.User
			fm.Name = g.Message.FunctionCall.Name

			var args map[string]string
			if err := json.Unmarshal([]byte(g.Message.FunctionCall.Arguments), &args); err != nil {
				fm.SetContent(fmt.Sprintf("error: %s", err.Error()))

				if err := fm.Save(ctx); err != nil {
					return common.ChatResult{}, err
				}

				msgs = append(msgs, messages.Message{
					ID:      fm.ID,
					Name:    fm.Name,
					Role:    messages.Role(fm.Role),
					Content: fm.Content,
				})

				opts = opts.WithMessages(msgs)
				return c.ChatWithOptions(ctx, opts)
			}

			req.Params = args

			if err := actions.HandleAction(req); (err != nil && !errors.Is(err, errors.ErrProceed)) || len(w.GetObjects()) == 0 {
				if err != nil {
					fm.SetContent(fmt.Sprintf("error: %s", err.Error()))
				} else {
					fm.SetContent("error: empty response")
				}

				if err := fm.Save(ctx); err != nil {
					return common.ChatResult{}, err
				}

				msgs = append(msgs, messages.Message{
					ID:      fm.ID,
					Name:    fm.Name,
					Role:    messages.Role(fm.Role),
					Content: fm.Content,
				})

				opts = opts.WithMessages(msgs)
				return c.ChatWithOptions(ctx, opts)
			}

			for ko := range w.GetObjects() {
				ffm := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "internal")
				ffm.Context = ctx
				ffm.Role = string(messages.Function)
				ffm.User = c.User
				ffm.Name = g.Message.FunctionCall.Name

				ffm.SetContent(string(w.GetObjects()[ko].Bytes()))

				if err := ffm.Save(ctx); err != nil {
					return common.ChatResult{}, err
				}

				msgs = append(msgs, messages.Message{
					ID:      ffm.ID,
					Name:    ffm.Name,
					Role:    messages.Role(ffm.Role),
					Content: ffm.Content,
				})
			}

			opts = opts.WithMessages(msgs)
			return c.ChatWithOptions(ctx, opts)
		}

	}

	res.Usage = common.ChatUsage{
		PromptTokens:     resp.Usage.PromptTokens,
		CompletionTokens: resp.Usage.CompletionTokens,
		TotalTokens:      resp.Usage.TotalTokens,
	}

	res.CalculateCosts(c.Model)

	logger.Info().
		Str("provider", "openai").
		Str("model", c.Model.Name).
		Float64("cost", res.Cost.TotalCost).
		Int("tokens", res.Usage.TotalTokens).
		Msg("chat completion done")

	return res, nil
}
