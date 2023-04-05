package tts

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/interfaces/discord"
	"github.com/kamushadenes/chloe/interfaces/slack"
	"github.com/kamushadenes/chloe/interfaces/telegram"
	"github.com/kamushadenes/chloe/memory"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type TTSAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewTTSAction() structs2.Action {
	return &TTSAction{
		Name: "tts",
	}
}

func (a *TTSAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *TTSAction) GetWriters() []io.WriteCloser {
	return a.Writers
}

func (a *TTSAction) GetName() string {
	return a.Name
}

func (a *TTSAction) GetNotification() string {
	return fmt.Sprintf("🔉 Generating audio: **%s**", a.Params)
}

func (a *TTSAction) SetParams(params string) {
	a.Params = params
}

func (a *TTSAction) GetParams() string {
	return a.Params
}

func (a *TTSAction) SetMessage(message *memory.Message) {}

func (a *TTSAction) Execute(request *structs.ActionRequest) error {
	errorCh := make(chan error)

	req := structs.NewTTSRequest()
	req.Context = request.GetContext()
	req.Content = a.Params
	req.ErrorChannel = errorCh

	req.Writers = a.Writers

	channels.TTSRequestsCh <- req

	for {
		select {
		case err := <-errorCh:
			if err != nil {
				return errors.Wrap(errors.ErrActionFailed, err)
			}
			return nil
		}
	}
}

func (a *TTSAction) RunPreActions(request *structs.ActionRequest) error {
	var ws []io.WriteCloser

	switch request.Message.Interface {
	case "telegram":
		iw := request.GetWriters()[0].(*telegram.TelegramWriter)
		aw := iw.ToAudioWriter()
		aw.SetPrompt(request.Params)
		ws = append(ws, aw)
	case "discord":
		iw := request.GetWriters()[0].(*discord.DiscordWriter)
		aw := iw.ToAudioWriter()
		aw.SetPrompt(request.Params)
		ws = append(ws, aw)
	case "slack":
		iw := request.GetWriters()[0].(*slack.SlackWriter)
		aw := iw.ToAudioWriter()
		aw.SetPrompt(request.Params)
		ws = append(ws, aw)
	default:
		ws = append(ws, request.GetWriters()[0])
	}

	a.SetWriters(ws)

	return nil
}

func (a *TTSAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
