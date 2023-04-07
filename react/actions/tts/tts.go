package tts

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
)

type TTSAction struct {
	Name   string
	Params map[string]string
}

func (a *TTSAction) GetNotification() string {
	return fmt.Sprintf("🔉 Generating audio: **%s**", a.Params["text"])
}

func (a *TTSAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Audio)

	errorCh := make(chan error)

	req := structs.NewTTSRequest()
	req.Context = request.GetContext()
	req.Content = a.Params["text"]
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
