package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
)

func processText(ctx context.Context, msg *memory.Message, ch chan interface{}) error {
	err := aiComplete(ctx, msg, ch)
	if err != nil {
		return err
	}

	return nil
}

func processAudio(ctx context.Context, msg *memory.Message) error {
	return aiTranscribe(ctx, msg)
}

func processImage(ctx context.Context, msg *memory.Message) error {
	return aiImage(ctx, msg)
}
