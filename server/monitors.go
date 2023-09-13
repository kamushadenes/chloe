package server

import (
	"context"
)

func MonitorMessages(ctx context.Context) {
	//logger := logging.GetLogger()

	/*for msg := range structs.IncomingMessagesCh {
		if err := ProcessMessage(ctx, msg); err != nil {
			logger.Err(err).Msg("failed to process message")
		}
	}*/
}

func MonitorRequests(ctx context.Context) {
	/*for {
		select {
		case req := <-structs.ActionRequestsCh:
			logger := structs.LoggerFromRequest(req)

			go func() {
				err := actions.HandleAction(req)
				if req.ErrorChannel != nil {
					req.ErrorChannel <- err
				}
				if err != nil {
					if errors.Is(err, errors.ErrProceed) {
						/*
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

							structs.CompletionRequestsCh <- req.ToCompletionRequest()
					} else {
						logger.Err(err).Msg("failed to handle action")
					}
				}
			}()
			/*
				case req := <-structs.CompletionRequestsCh:
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
				case req := <-structs.TranscribeRequestsCh:
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
				case req := <-structs.GenerationRequestsCh:
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
				case req := <-structs.EditRequestsCh:
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
				case req := <-structs.VariationRequestsCh:
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
				case req := <-structs.TTSRequestsCh:
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
	*/
}
