package slack

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
)

func complete(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	if err := aiComplete(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating image")
	}
}

func generate(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	if err := aiGenerate(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating image")
	}
}

func tts(ctx context.Context, msg *memory.Message) {
	logger := zerolog.Ctx(ctx)

	if err := aiTTS(ctx, msg); err != nil {
		logger.Error().Err(err).Msg("error generating audio")
	}
}

func forgetUser(ctx context.Context, msg *memory.Message) error {
	return msg.User.DeleteMessages(ctx)
}
