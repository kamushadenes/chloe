package openai

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
)

// newMessageModeration creates a new memory.MessageModeration struct
// from the provided openai.ModerationResponse.
func newMessageModeration(result openai.ModerationResponse) *memory.MessageModeration {
	r := result.Results[0]
	return &memory.MessageModeration{
		CategoryHate:                 r.Categories.Hate,
		CategoryHateThreatening:      r.Categories.HateThreatening,
		CategorySelfHarm:             r.Categories.SelfHarm,
		CategorySexual:               r.Categories.Sexual,
		CategorySexualMinors:         r.Categories.SexualMinors,
		CategoryViolence:             r.Categories.Violence,
		CategoryViolenceGraphic:      r.Categories.ViolenceGraphic,
		CategoryScoreHate:            r.CategoryScores.Hate,
		CategoryScoreHateThreatening: r.CategoryScores.HateThreatening,
		CategoryScoreSelfHarm:        r.CategoryScores.SelfHarm,
		CategoryScoreSexual:          r.CategoryScores.Sexual,
		CategoryScoreSexualMinors:    r.CategoryScores.SexualMinors,
		CategoryScoreViolence:        r.CategoryScores.Violence,
		CategoryScoreViolenceGraphic: r.CategoryScores.ViolenceGraphic,
		Flagged:                      r.Flagged,
	}
}

// newModerationRequest creates a new openai.ModerationRequest from the provided memory.Message.
// Returns an error if the message has no content to moderate.
func newModerationRequest(msg *memory.Message) (openai.ModerationRequest, error) {
	model := config.OpenAI.DefaultModel.Moderation
	content := msg.Content

	if content == "" {
		return openai.ModerationRequest{}, fmt.Errorf("no content to moderate")
	}

	req := openai.ModerationRequest{
		Input: content,
		Model: &model,
	}

	return req, nil
}

// createModerationWithTimeout attempts to create a ModerationResponse with a timeout.
// Returns the created ModerationResponse or an error if the request times out or encounters an issue.
func createModerationWithTimeout(ctx context.Context, req openai.ModerationRequest) (openai.ModerationResponse, error) {
	logger := zerolog.Ctx(ctx)

	respi, err := utils.WaitTimeout(ctx, config.Timeouts.Moderation, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.Moderations(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error moderating message")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return openai.ModerationResponse{}, err
	}

	return respi.(openai.ModerationResponse), nil
}

// Moderate processes a moderation request for a message using the OpenAI API.
// Returns an error if there's an issue during the process.
func Moderate(ctx context.Context, msg *memory.Message) error {
	logger := logging.GetLogger().With().Uint("messageId", msg.ID).Logger()

	logger.Info().Msg("moderating message")

	req, err := newModerationRequest(msg)
	if err != nil {
		return err
	}

	resp, err := createModerationWithTimeout(ctx, req)
	if err != nil {
		return err
	}

	msg.Moderation = newMessageModeration(resp)
	msg.Moderated = true

	return msg.Save(ctx)
}
