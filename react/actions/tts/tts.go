package tts

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

type TTSAction struct {
	Name   string
	Params string
}

func NewTTSAction() structs.Action {
	return &TTSAction{
		Name: "tts",
	}
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

func (a *TTSAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Audio)

	errorCh := make(chan error)

	req := structs.NewTTSRequest()
	req.Context = request.GetContext()
	req.Content = a.Params
	req.ErrorChannel = errorCh

	req.Writer = obj

	channels.TTSRequestsCh <- req

	for {
		select {
		case err := <-errorCh:
			if err != nil {
				return nil, errors.Wrap(errors.ErrActionFailed, err)
			}
			return []*structs.ResponseObject{obj}, nil
		}
	}
}

func (a *TTSAction) RunPreActions(request *structs.ActionRequest) error {
	return nil
}

func (a *TTSAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
