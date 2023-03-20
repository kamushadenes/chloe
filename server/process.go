package server

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/messages"
	"github.com/rs/zerolog"
	"os"
)

func ProcessMessage(ctx context.Context, msg *messages.Message) error {
	logger := zerolog.Ctx(ctx)

	logger.Info().Str("userId", msg.User.ID).Str("interface", msg.User.ExternalID.Interface).Msg("message received")

	return nil
}

func DeliverMessage(ctx context.Context, msg *channels.OutgoingMessage) error {
	logger := zerolog.Ctx(ctx)

	switch msg.User.ExternalID.Interface {
	case "telegram":
		for k := range msg.Texts {
			text := msg.Texts[k]
			for kk := range msg.TextWriters {
				writer := msg.TextWriters[kk]
				writer.Write([]byte(text))
			}
		}

		for k := range msg.Audios {
			audio := msg.Audios[k]
			for kk := range msg.AudioWriters {
				writer := msg.AudioWriters[kk]

				b, err := os.ReadFile(audio)
				if err != nil {
					logger.Err(err).Msg("failed to open audio file")
					continue
				}
				writer.Write(b)
			}
		}

		for k := range msg.Images {
			image := msg.Images[k]
			for kk := range msg.ImageWriters {
				writer := msg.ImageWriters[kk]

				b, err := os.ReadFile(image)
				if err != nil {
					logger.Err(err).Msg("failed to open image file")
					continue
				}
				writer.Write(b)
			}
		}
	}

	return nil
}
