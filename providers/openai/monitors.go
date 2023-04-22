package openai

import (
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
	"time"
)

func MonitorModeration(ctx context.Context) {
	logger := logging.GetLogger()
	ticker := time.NewTicker(30 * time.Second)

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

func MonitorSummary(ctx context.Context) {
	logger := logging.GetLogger()
	ticker := time.NewTicker(30 * time.Second)

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
