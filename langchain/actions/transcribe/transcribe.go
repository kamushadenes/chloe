package transcribe

import (
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *TranscribeAction) GetNotification() string {
	return "✍️ Transcribing..."
}

func (a *TranscribeAction) SetMessage(message *memory.Message) {
	a.Extra["message"] = message
}

func (a *TranscribeAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	/*
		obj := structs.NewResponseObject(structs.Text)

		errorCh := make(chan error)

		req := structs.NewTranscriptionRequest()
		req.Context = request.GetContext()
		req.Message = a.Message
		req.FilePath = a.Params["path"]
		req.ErrorChannel = errorCh
		req.Writer = obj

		//structs.TranscribeRequestsCh <- req

		for {
			select {
			case err := <-errorCh:
				if err != nil {
					return nil, errors.Wrap(errors.ErrActionFailed, err)
				}
				return []*response_object_structs.ResponseObject{obj}, nil
			}
		}
	*/

	return nil, nil
}

func (a *TranscribeAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrProceed
}
