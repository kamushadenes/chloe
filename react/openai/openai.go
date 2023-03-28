package utils

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/resources"
	"github.com/kamushadenes/chloe/timeout"
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

	respi, err := timeout.WaitTimeout(ctx, config.Timeouts.ChainOfThought, func(ch chan interface{}, errCh chan error) {
		resp, err := OpenAIClient.CreateChatCompletion(ctx, req)
		if err != nil {
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return openai.ChatCompletionResponse{}, err
	}

	return respi.(openai.ChatCompletionResponse), nil
}
