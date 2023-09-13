package slack

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/slack-go/slack"
)

func (w *SlackWriter) closeText() error {
	logger := logging.FromContext(w.Context)

	for kk := range w.objs {
		obj := w.objs[kk]
		if obj.Type == response_object_structs.Text {
			logger.Debug().Str("chatID", w.ChatID).Msg("replying with text")
			msgs := utils.StringToChunks(obj.String(), config.Slack.MaxMessageLength)

			if w.externalID == "" {
				for k := range msgs {
					if err := w.Request.GetMessage().SendText(msgs[k], true); err != nil {
						return err
					}
				}
			} else {
				if _, _, _, err := w.Bot.UpdateMessage(w.ChatID, w.externalID, slack.MsgOptionText(msgs[0], false)); err != nil {
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
