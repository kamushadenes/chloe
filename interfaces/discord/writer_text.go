package discord

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
)

func (w *DiscordWriter) closeText() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	logger.Debug().Str("chatID", w.ChatID).Msg("replying with text")

	for kk := range w.objs {
		obj := w.objs[kk]
		if obj.Type == structs.Text {
			msgs := utils.StringToChunks(obj.String(), config.Discord.MaxMessageLength)

			if w.externalID == "" {
				for k := range msgs {
					if err := w.Request.GetMessage().SendText(msgs[k], true); err != nil {
						return err
					}
				}
			} else {
				if _, err := w.Bot.ChannelMessageEdit(w.ChatID, w.externalID, msgs[0]); err != nil {
					return err
				}

				for k := range msgs[1:] {
					if err := w.Request.GetMessage().SendText(msgs[k], true); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
