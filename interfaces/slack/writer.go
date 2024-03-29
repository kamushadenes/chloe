package slack

import (
	"context"
	"github.com/kamushadenes/chloe/langchain/memory"
	"net/http"
	"time"

	"github.com/aquilax/truncate"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
	"github.com/slack-go/slack"
)

type SlackWriter struct {
	Context          context.Context
	Prompt           string
	Bot              *slack.Client
	ChatID           string
	Type             string
	ReplyID          string
	Message          *memory.Message
	objs             []*response_object_structs.ResponseObject
	externalID       string
	lastUpdate       *time.Time
	preWriteCallback func()
}

func NewSlackWriter(ctx context.Context, msg *memory.Message, reply bool, prompt ...string) *SlackWriter {
	var chatID string
	var ts string

	if msg.Source.Slack.Message != nil {
		chatID = msg.Source.Slack.Message.Channel
		ts = msg.Source.Slack.Message.TimeStamp
	} else {
		chatID = msg.Source.Slack.SlashCommand.ChannelID
	}

	w := &SlackWriter{
		Context: ctx,
		Bot:     msg.Source.Slack.API,
		ChatID:  chatID,
		Message: msg,
	}

	if len(prompt) > 0 {
		w.Prompt = prompt[0]
	}

	if reply && ts != "" {
		w.ReplyID = ts
	}

	return w
}

func (w *SlackWriter) Flush() {
	if w.Type != "text" {
		return
	}

	if w.externalID == "" && (config.Slack.SendProcessingMessage || config.Slack.StreamMessages) {
		_, ts, err := w.Bot.PostMessage(w.ChatID, slack.MsgOptionText("↻ Processing...", false))
		if err != nil {
			return
		}
		w.externalID = ts
		tt := time.Now()
		w.lastUpdate = &tt
	}

	if !config.Slack.StreamMessages {
		return
	}

	if w.lastUpdate != nil && time.Since(*w.lastUpdate) > config.Slack.StreamFlushInterval {
		_, _, _, _ = w.Bot.UpdateMessage(w.ChatID, w.externalID, slack.MsgOptionText(
			truncate.Truncate(
				w.objs[0].String(),
				config.Slack.MaxMessageLength,
				"...",
				truncate.PositionEnd,
			), false))

		tt := time.Now()
		w.lastUpdate = &tt
	}
}

func (w *SlackWriter) Close() error {
	funcs := []func() error{
		w.closeText,
		w.closeAudio,
		w.closeImage,
	}

	for k := range funcs {
		if err := funcs[k](); err != nil {
			return err
		}
	}

	return nil
}

func (w *SlackWriter) Write(p []byte) (n int, err error) {
	if w.preWriteCallback != nil {
		w.preWriteCallback()
	}
	if len(w.objs) == 0 {
		w.objs = append(w.objs, &response_object_structs.ResponseObject{
			Type:   response_object_structs.Text,
			Result: true,
		})
	}

	w.objs[0].Data = append(w.objs[0].Data, p...)

	return len(p), nil
}

func (w *SlackWriter) WriteObject(obj *response_object_structs.ResponseObject) error {
	w.objs = append(w.objs, obj)

	return nil
}

func (w *SlackWriter) SetPrompt(prompt string) {
	w.Prompt = prompt
}

func (w *SlackWriter) WriteHeader(int)     {}
func (w *SlackWriter) Header() http.Header { return http.Header{} }
func (w *SlackWriter) SetPreWriteCallback(fn func()) {
	w.preWriteCallback = fn
}

func (w *SlackWriter) GetObjects() []*response_object_structs.ResponseObject {
	return w.objs
}
