package transcribe

import (
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
)

type TranscribeAction struct {
	Name    string
	Params  map[string]string
	Message *memory.Message
}

func (a *TranscribeAction) GetNotification() string {
	return "✍️ Transcribing..."
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
	req.FilePath = a.Params["path"]
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

func (a *TranscribeAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrProceed
}
