package openai

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/resources"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"strings"
	"time"
)

func Summarize(ctx context.Context, msg *memory.Message) error {
	logger := zerolog.Ctx(ctx).With().Uint("messageId", msg.ID).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("summarizing text")

	promptSize, err := resources.GetPromptSize("summarize")
	if err != nil {
		return err
	}

	prompt, err := resources.GetPrompt("summarize", &resources.PromptArgs{
		Args: map[string]interface{}{
			"text": react.Truncate(msg.Content,
				int(float64(config.OpenAI.MaxTokens[config.OpenAI.DefaultModel.Summarization])-
					(float64(promptSize)*0.75)-
					(float64(len(strings.Fields(msg.Content)))*0.75))),
		},
		Mode: "summarize",
	})
	if err != nil {
		return err
	}

	req := openai.ChatCompletionRequest{
		Model: config.OpenAI.DefaultModel.Summarization,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
		},
	}

	var response openai.ChatCompletionResponse

	respi, err := utils.WaitTimeout(ctx, config.Timeouts.Completion, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateChatCompletion(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error summarizing message")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return err
	}

	response = respi.(openai.ChatCompletionResponse)

	return msg.SetSummary(ctx, response.Choices[0].Message.Content)
}

func MonitorSummary(ctx context.Context) {
	logger := zerolog.Ctx(ctx)
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			messages, err := memory.LoadNonSummarizedMessages(ctx)
			if err != nil {
				logger.Error().Err(err).Msg("failed to load non summarized messages")
				continue
			}
			for k := range messages {
				err := Summarize(ctx, messages[k])
				if err != nil {
					logger.Error().Err(err).Msg("failed to summarize message")
					continue
				}
			}
		}
	}
}
