package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
)

func (w *TelegramWriter) closeImage() error {
	logger := logging.FromContext(w.Context)

	var files []interface{}
	for k := range w.objs {
		obj := w.objs[k]
		if obj.Type == structs.Image {
			files = append(files, tgbotapi.NewInputMediaPhoto(
				tgbotapi.FileReader{
					Name:   obj.Name,
					Reader: obj,
				},
			))
		}
	}

	if len(files) == 0 {
		return nil
	}

	logger.Debug().Int64("chatID", w.ChatID).Msg("replying with image")

	msg := tgbotapi.NewMediaGroup(w.ChatID, files)
	msg.ReplyToMessageID = w.ReplyID
	_, err := w.Bot.SendMediaGroup(msg)
	return err
}
