package telegram

import (
	"context"
	"github.com/kamushadenes/chloe/messages"
)

func processText(ctx context.Context, msg *messages.Message, ch chan interface{}) error {
	err := aiComplete(ctx, msg, ch)
	if err != nil {
		return err
	}

	return nil
}

func processAudio(ctx context.Context, msg *messages.Message, ch chan interface{}) error {
	return aiTranscribe(ctx, msg, ch)
}

func processImage(ctx context.Context, msg *messages.Message) error {
	return aiImage(ctx, msg)
}
