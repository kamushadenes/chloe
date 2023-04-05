package transcribe

import (
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

type TranscribeAction struct {
	Name    string
	Params  string
	Message *memory.Message
}

func NewTranscribeAction() structs.Action {
	return &TranscribeAction{
		Name: "audio",
	}
}

func (a *TranscribeAction) GetName() string {
	return a.Name
}

func (a *TranscribeAction) GetNotification() string {
	return "✍️ Transcribing..."
}

func (a *TranscribeAction) SetParams(params string) {
	a.Params = params
}

func (a *TranscribeAction) GetParams() string {
	return a.Params
}

func (a *TranscribeAction) SetMessage(message *memory.Message) {
	a.Message = message
}

func (a *TranscribeAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	errorCh := make(chan error)

	req := structs.NewTranscriptionRequest()
	req.Context = request.GetContext()
	req.Message = a.Message
	req.FilePath = a.Params
	req.ErrorChannel = errorCh
	req.Writer = obj

	channels.TranscribeRequestsCh <- req

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

func (a *TranscribeAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *TranscribeAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}
