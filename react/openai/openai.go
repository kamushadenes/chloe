package utils

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/cost"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/prompts"
	"github.com/kamushadenes/chloe/timeouts"
	"github.com/sashabaranov/go-openai"
)

var OpenAIClient *openai.Client

func init() {
	if config.OpenAI.UseAzure {
		OpenAIClient = openai.NewClientWithConfig(openai.DefaultAzureConfig(config.OpenAI.APIKey, config.OpenAI.AzureBaseURL, config.OpenAI.AzureEngine))
	} else {
		OpenAIClient = openai.NewClient(config.OpenAI.APIKey)
	}
}

func SimpleCompletionRequest(ctx context.Context, prompt string, message string) (openai.ChatCompletionResponse, error) {
	prompt, err := prompts.GetPrompt(prompt, &prompts.PromptArgs{
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

	promptPrice, promptUnitSize, completionPrice, completionUnitSize := config.OpenAI.GetModelCostInfo(config.Completion)

	promptCost := promptPrice * float64(respi.(openai.ChatCompletionResponse).Usage.PromptTokens) / float64(promptUnitSize)
	completionCost := completionPrice * float64(respi.(openai.ChatCompletionResponse).Usage.CompletionTokens) / float64(completionUnitSize)

	cost.AddCategoryCost(string(config.ChainOfThought), promptCost+completionCost)

	return respi.(openai.ChatCompletionResponse), nil
}
