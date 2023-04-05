package utils

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/resources"
	"github.com/kamushadenes/chloe/timeouts"
	"github.com/sashabaranov/go-openai"
)

var OpenAIClient = openai.NewClient(config.OpenAI.APIKey)

func SimpleCompletionRequest(ctx context.Context, prompt string, message string) (openai.ChatCompletionResponse, error) {
	prompt, err := resources.GetPrompt(prompt, &resources.PromptArgs{
		Args: map[string]interface{}{},
		Mode: prompt,
	})

	req := openai.ChatCompletionRequest{
		Model: config.OpenAI.DefaultModel.ChainOfThought.String(),
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
			{
				Role:    "user",
				Content: message,
			},
		},
	}

	respi, err := timeouts.WaitTimeout(ctx, config.Timeouts.ChainOfThought, func(ch chan interface{}, errCh chan error) {
		resp, err := OpenAIClient.CreateChatCompletion(ctx, req)
		if err != nil {
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return openai.ChatCompletionResponse{}, errors.Wrap(errors.ErrCompletionFailed, err)
	}

	return respi.(openai.ChatCompletionResponse), nil
}
