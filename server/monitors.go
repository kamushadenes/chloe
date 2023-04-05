package server

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/providers/google"
	"github.com/kamushadenes/chloe/providers/openai"
	"github.com/kamushadenes/chloe/react"
	"github.com/kamushadenes/chloe/react/actions"
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
		case req := <-channels.ActionRequestsCh:
			logger := structs.LoggerFromRequest(req)
			go func() {
				err := actions.HandleAction(req)
				if req.ErrorChannel != nil {
					req.ErrorChannel <- err
				}
				if err != nil {
					if errors.Is(err, errors.ErrProceed) {
						// duplicate code
						msg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), req.Message.Interface)
						msg.SetContent(fmt.Sprintf("summarize everything since \"%s\", never mention the checkpoint, try to make it look like a news article or a Wikipedia page, don't provide explanation", react.CheckpointMarker))
						msg.ErrorCh = req.Message.ErrorCh
						msg.Role = "user"
						msg.User = req.Message.User
						msg.Context = req.Message.Context
						msg.Source = req.Message.Source
						if err := msg.Save(req.GetContext()); err != nil {
							_ = msg.SendError(errors.Wrap(errors.ErrSaveMessage, err))
						}

						channels.CompletionRequestsCh <- req.ToCompletionRequest()
					} else {
						logger.Err(err).Msg("failed to handle action")
					}
				}
			}()
		case req := <-channels.CompletionRequestsCh:
			logger := structs.LoggerFromRequest(req)
			go func() {
				err := openai.Complete(req)
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
				err := openai.Transcribe(req)
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
				err := openai.Generate(req)
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
				err := openai.Edits(req)
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
				err := openai.Variations(req)
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
				err := google.TTS(req)
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
