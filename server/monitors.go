package server

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/providers/google"
	"github.com/kamushadenes/chloe/providers/openai"
	"github.com/rs/zerolog"
)

func MonitorMessages(ctx context.Context) {
	logger := zerolog.Ctx(ctx)

	for {
		select {
		case msg := <-channels.IncomingMessagesCh:
			if err := ProcessMessage(ctx, msg); err != nil {
				logger.Err(err).Msg("failed to process message")
			}
		case msg := <-channels.OutgoingMessagesCh:
			if err := DeliverMessage(ctx, msg); err != nil {
				logger.Err(err).Msg("failed to deliver message")
			}
		}
	}
}

func MonitorRequests(ctx context.Context) {
	logger := zerolog.Ctx(ctx)

	for {
		select {
		case req := <-channels.CompletionRequestsCh:
			go func() {
				err := openai.Complete(req.Context, req)
				if req.ErrorChannel != nil {
					req.ErrorChannel <- err
				}
				if err != nil {
					logger.Err(err).Msg("failed to complete text")
				}
			}()
		case req := <-channels.TranscribeRequestsCh:
			go func() {
				err := openai.Transcribe(req.Context, req)
				if req.ErrorChannel != nil {
					req.ErrorChannel <- err
				}
				if err != nil {
					logger.Err(err).Msg("failed to transcribe audio")
				}
			}()
		case req := <-channels.GenerationRequestsCh:
			go func() {
				err := openai.Generate(req.Context, req)
				if req.ErrorChannel != nil {
					req.ErrorChannel <- err
				}
				if err != nil {
					logger.Err(err).Msg("failed to generate image")
				}
			}()
		case req := <-channels.EditRequestsCh:
			go func() {
				err := openai.Edits(req.Context, req)
				if req.ErrorChannel != nil {
					req.ErrorChannel <- err
				}
				if err != nil {
					logger.Err(err).Msg("failed to edit image")
				}
			}()
		case req := <-channels.VariationRequestsCh:
			go func() {
				err := openai.Variations(req.Context, req)
				if req.ErrorChannel != nil {
					req.ErrorChannel <- err
				}
				if err != nil {
					logger.Err(err).Msg("failed to create image variations")
				}
			}()
		case req := <-channels.TTSRequestsCh:
			go func() {
				err := google.TTS(req.Context, req)
				if req.ErrorChannel != nil {
					req.ErrorChannel <- err
				}
				if err != nil {
					logger.Err(err).Msg("failed to generate audio")
				}
			}()
		}
	}
}
