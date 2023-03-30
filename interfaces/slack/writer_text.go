package slack

import (
	"bytes"
	"context"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
)

func (w *SlackWriter) ToTextWriter() *SlackWriter {
	return NewTextWriter(w.Context, w.Request, w.ReplyID != "", w.Prompt)
}

func NewTextWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *SlackWriter {
	var chatID string
	var ts string

	if request.GetMessage().Source.Slack.Message != nil {
		chatID = request.GetMessage().Source.Slack.Message.Channel
		ts = request.GetMessage().Source.Slack.Message.TimeStamp
	} else {
		chatID = request.GetMessage().Source.Slack.SlashCommand.ChannelID
	}

	w := &SlackWriter{
		Context: ctx,
		Bot:     request.GetMessage().Source.Slack.API,
		ChatID:  chatID,
		Type:    "text",
		Request: request,
		bufs:    []bytes.Buffer{{}},
		bufID:   0,
	}

	if reply && ts != "" {
		w.ReplyID = ts
	}
	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	return w
}

func (w *SlackWriter) closeText() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	logger.Debug().Str("chatID", w.ChatID).Msg("replying with text")

	msgs := utils.StringToChunks(w.bufs[0].String(), config.Slack.MaxMessageLength)

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

	return nil
}
