package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
)

func processText(ctx context.Context, msg *memory.Message) error {
	return aiComplete(ctx, msg)
}

func processAudio(ctx context.Context, msg *memory.Message) error {
	return aiTranscribe(ctx, msg)
}

func processImage(ctx context.Context, msg *memory.Message) error {
	return aiImage(ctx, msg)
}
