package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
	"github.com/kamushadenes/chloe/utils"
)

func (w *TelegramWriter) closeText() error {
	logger := logging.FromContext(w.Context)

	for kk := range w.objs {
		obj := w.objs[kk]
		if obj.Type == response_object_structs.Text {
			logger.Debug().Int64("chatID", w.ChatID).Msg("replying with text")

			msgs := utils.StringToChunks(obj.String(), config.Discord.MaxMessageLength)

			if w.externalID == 0 {
				for k := range msgs {
					if err := w.Message.SendText(msgs[k], true, w.ReplyID); err != nil {
						return err
					}
				}
			} else {
				if _, err := w.Bot.Send(tgbotapi.NewEditMessageText(w.ChatID, w.externalID, msgs[0])); err != nil {
					return err
				}

				for k := range msgs[1:] {
					if err := w.Message.SendText(msgs[k], true); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}
