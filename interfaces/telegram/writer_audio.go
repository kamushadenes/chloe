package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (w *TelegramWriter) closeAudio() error {
	logger := logging.FromContext(w.Context)

	for k := range w.objs {
		obj := w.objs[k]
		if obj.Type == response_object_structs.Audio {
			logger.Debug().Int64("chatID", w.ChatID).Msg("replying with audio")

			tmsg := tgbotapi.NewVoice(w.ChatID, tgbotapi.FileReader{
				Name:   obj.Name,
				Reader: obj,
			})

			tmsg.ReplyToMessageID = w.ReplyID
			_, err := w.Bot.Send(tmsg)
			if err != nil {
				logger.Error().Err(err).Msg("failed to send audio")
			}
		}
	}

	return nil
}
