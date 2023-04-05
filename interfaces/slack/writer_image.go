package slack

import (
	"fmt"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
	"github.com/slack-go/slack"
)

func (w *SlackWriter) closeImage() error {
	logger := logging.FromContext(w.Context)

	content := fmt.Sprintf("Prompt: %s", w.Prompt)

	_, ts, err := w.Bot.PostMessage(w.ChatID, slack.MsgOptionText(content, false))
	if err != nil {
		return err
	}

	for k := range w.objs {
		obj := w.objs[k]
		if obj.Type == structs.Image {
			logger.Debug().Str("chatID", w.ChatID).Msg("replying with image")
			if _, err := w.Bot.UploadFileV2(slack.UploadFileV2Parameters{
				Reader:          obj,
				FileSize:        obj.Size(),
				Filename:        fmt.Sprintf("generated-%s-%d.png", ts, k),
				Title:           content,
				Channel:         w.ChatID,
				AltTxt:          content,
				ThreadTimestamp: ts,
			}); err != nil {
				logger.Error().Err(err).Msg("failed to upload image")
			}
		}

	}

	return nil
}
