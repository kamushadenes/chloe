package slack

import (
	"bytes"
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"github.com/slack-go/slack"
)

func (w *SlackWriter) ToImageWriter() *SlackWriter {
	return NewImageWriter(w.Context, w.Request, w.ReplyID != "", w.Prompt)
}

func NewImageWriter(ctx context.Context, request structs.ActionOrCompletionRequest, reply bool, prompt ...string) *SlackWriter {
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
		Type:    "image",
		Request: request,
		bufs:    []bytes.Buffer{},
		bufID:   0,
	}
	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	if reply && ts != "" {
		w.ReplyID = ts
	}

	return w
}

func (w *SlackWriter) closeImage() error {
	logger := zerolog.Ctx(w.Context).With().Str("requestID", w.Request.GetID()).Logger()

	bufs := w.bufs

	if w.mainWriter != nil {
		bufs = w.mainWriter.bufs
		w.mainWriter.closedBufs++
		if w.mainWriter.closedBufs != len(w.mainWriter.bufs) {
			return nil
		}
	}

	logger.Debug().Str("chatID", w.ChatID).Msg("replying with image")

	content := fmt.Sprintf("Prompt: %s", w.Prompt)

	_, ts, err := w.Bot.PostMessage(w.ChatID, slack.MsgOptionText(content, false))
	if err != nil {
		return err
	}

	for k := range bufs {
		if _, err := w.Bot.UploadFileV2(slack.UploadFileV2Parameters{
			Reader:          bytes.NewReader(bufs[k].Bytes()),
			FileSize:        len(bufs[k].Bytes()),
			Filename:        fmt.Sprintf("generated-%s-%d.png", ts, k),
			Title:           content,
			Channel:         w.ChatID,
			AltTxt:          content,
			ThreadTimestamp: ts,
		}); err != nil {
			return err
		}
	}

	return nil
}
