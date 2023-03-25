package server

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/providers/google"
	"github.com/kamushadenes/chloe/providers/openai"
	"github.com/kamushadenes/chloe/structs"
)

func MonitorMessages(ctx context.Context) {
	logger := logging.GetLogger()

	for {
		select {
		case msg := <-channels.IncomingMessagesCh:
			if err := ProcessMessage(ctx, msg); err != nil {
				logger.Err(err).Msg("failed to process message")
			}
		}
	}
}

func MonitorRequests(ctx context.Context) {
	for {
		select {
		case req := <-channels.CompletionRequestsCh:
			logger := structs.LoggerFromRequest(req)
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
			logger := structs.LoggerFromRequest(req)
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
			logger := structs.LoggerFromRequest(req)
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
			logger := structs.LoggerFromRequest(req)
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
			logger := structs.LoggerFromRequest(req)
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
			logger := structs.LoggerFromRequest(req)
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
