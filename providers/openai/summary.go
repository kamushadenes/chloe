package openai

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	utils2 "github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/resources"
	"github.com/kamushadenes/chloe/timeout"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"strings"
)

// getSummarizationPrompt retrieves a summarization prompt for the given message.
// Returns an error if there's an issue during the process.
func getSummarizationPrompt(ctx context.Context, msg *memory.Message) (string, error) {
	promptSize, err := resources.GetPromptSize("summarize")
	if err != nil {
		return "", err
	}

	return resources.GetPrompt("summarize", &resources.PromptArgs{
		Args: map[string]interface{}{
			"text": utils2.Truncate(msg.Content,
				int(float64(config.OpenAI.MaxTokens[config.OpenAI.DefaultModel.Summarization])-
					(float64(promptSize)*0.75)-
					(float64(len(strings.Fields(msg.Content)))*0.75))),
		},
		Mode: "summarize",
	})
}

// newSummarizationRequest creates a new openai.ChatCompletionRequest for summarization.
// Returns an error if there's an issue during the process.
func newSummarizationRequest(ctx context.Context, msg *memory.Message) (openai.ChatCompletionRequest, error) {
	prompt, err := getSummarizationPrompt(ctx, msg)
	if err != nil {
		return openai.ChatCompletionRequest{}, err
	}

	return openai.ChatCompletionRequest{
		Model: config.OpenAI.DefaultModel.Summarization,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "system",
				Content: prompt,
			},
		},
	}, nil
}

// createSummarizationWithTimeout attempts to create a ChatCompletionResponse with a timeout.
// Returns the created ChatCompletionResponse or an error if the request times out or encounters an issue.
func createSummarizationWithTimeout(ctx context.Context, req openai.ChatCompletionRequest) (openai.ChatCompletionResponse, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := timeout.WaitTimeout(ctx, config.Timeouts.Completion, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.CreateChatCompletion(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error summarizing message")
			errCh <- err
		}
		ch <- resp
	})

	return respi.(openai.ChatCompletionResponse), err
}

// Summarize processes a summarization request for a message using the OpenAI API.
// Returns an error if there's an issue during the process.
func Summarize(ctx context.Context, msg *memory.Message) error {
	logger := logging.GetLogger().With().Uint("messageId", msg.ID).Logger()
	ctx = logger.WithContext(ctx)

	logger.Info().Msg("summarizing text")

	req, err := newSummarizationRequest(ctx, msg)
	if err != nil {
		return err
	}

	response, err := createSummarizationWithTimeout(ctx, req)
	if err != nil {
		return err
	}

	return msg.SetSummary(ctx, response.Choices[0].Message.Content)
}
