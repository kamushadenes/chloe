package openai

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/sashabaranov/go-openai"
	"time"
)

func Moderate(ctx context.Context, msg *memory.Message) error {
	logger := zerolog.Ctx(ctx).With().Uint("messageId", msg.ID).Logger()

	logger.Info().Msg("moderating message")

	model := config.OpenAI.DefaultModel[config.ModelPurposeModeration]
	content := msg.Content
	if content == "" {
		content = msg.ChainOfThought
	}

	if content == "" {
		return fmt.Errorf("no content to moderate")
	}

	req := openai.ModerationRequest{
		Input: content,
		Model: &model,
	}

	var resp openai.ModerationResponse

	respi, err := utils.WaitTimeout(ctx, config.TimeoutTypeModeration, func(ch chan interface{}, errCh chan error) {
		resp, err := openAIClient.Moderations(ctx, req)
		if err != nil {
			logger.Error().Err(err).Msg("error moderating message")
			errCh <- err
		}
		ch <- resp
	})
	if err != nil {
		return err
	}

	resp = respi.(openai.ModerationResponse)

	result := resp.Results[0]

	msg.Moderation = &memory.MessageModeration{
		CategoryHate:                 result.Categories.Hate,
		CategoryHateThreatening:      result.Categories.HateThreatening,
		CategorySelfHarm:             result.Categories.SelfHarm,
		CategorySexual:               result.Categories.Sexual,
		CategorySexualMinors:         result.Categories.SexualMinors,
		CategoryViolence:             result.Categories.Violence,
		CategoryViolenceGraphic:      result.Categories.ViolenceGraphic,
		CategoryScoreHate:            result.CategoryScores.Hate,
		CategoryScoreHateThreatening: result.CategoryScores.HateThreatening,
		CategoryScoreSelfHarm:        result.CategoryScores.SelfHarm,
		CategoryScoreSexual:          result.CategoryScores.Sexual,
		CategoryScoreSexualMinors:    result.CategoryScores.SexualMinors,
		CategoryScoreViolence:        result.CategoryScores.Violence,
		CategoryScoreViolenceGraphic: result.CategoryScores.ViolenceGraphic,
		Flagged:                      result.Flagged,
	}
	msg.Moderated = true

	return msg.Save(ctx)
}

func MonitorModeration(ctx context.Context) {
	logger := zerolog.Ctx(ctx)
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if config.OpenAI.ModerateMessages {
				messages, err := memory.LoadNonModeratedMessages(ctx)
				if err != nil {
					logger.Error().Err(err).Msg("failed to load non moderated messages")
					continue
				}
				for k := range messages {
					err := Moderate(ctx, messages[k])
					if err != nil {
						logger.Error().Err(err).Msg("failed to moderate message")
						continue
					}
				}
			}
		}
	}
}
