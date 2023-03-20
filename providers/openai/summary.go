package openai

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/resources"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"time"
)

func Summarize(ctx context.Context, id string, text string) (string, error) {
	logger := zerolog.Ctx(ctx).With().Str("messageId", id).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("summarizing text")

	promptSize, err := resources.GetPromptSize("summarize")
	if err != nil {
		return "", err
	}

	prompt, err := resources.GetPrompt("summarize", &resources.PromptArgs{
		Args: map[string]interface{}{
			"text": react.Truncate(text,
				int(float64(config.OpenAI.MaxTokens[config.OpenAI.DefaultModel[config.ModelPurposeCompletion]])-
					(float64(promptSize)*0.75)-
					(float64(len(text))*0.75))),
		},
		Mode: "summarize",
	})
	if err != nil {
		return "", err
	}

	req := openai.ChatCompletionRequest{
		Model: config.OpenAI.DefaultModel[config.ModelPurposeCompletion],
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
		},
	}

	response, err := openAIClient.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	return response.Choices[0].Message.Content, nil
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
				summary, err := Summarize(ctx, messages[k][0], messages[k][1])
				if err != nil {
					logger.Error().Err(err).Msg("failed to summarize message")
					continue
				}
				err = memory.SetMessageSummary(ctx, messages[k][0], summary)
				if err != nil {
					logger.Error().Err(err).Msg("failed to save message")
					continue
				}
			}
		}
	}
}
