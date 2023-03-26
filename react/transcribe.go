package react

import (
	"context"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type TranscribeAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
	User    *memory.User
	Message *memory.Message
}

func NewTranscribeAction() *TranscribeAction {
	return &TranscribeAction{
		Name: "audio",
	}
}

func (a *TranscribeAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *TranscribeAction) GetWriters() []io.WriteCloser {
	return a.Writers
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

func (a *TranscribeAction) SetUser(user *memory.User) {
	a.User = user
}
func (a *TranscribeAction) SetMessage(message *memory.Message) {
	a.Message = message
}

func (a *TranscribeAction) Execute(ctx context.Context) error {
	errorCh := make(chan error)

	req := structs.NewTranscriptionRequest()
	req.Context = ctx
	req.Message = a.Message
	req.FilePath = a.Params
	req.ErrorChannel = errorCh
	req.Writer = a.Writers[0]

	channels.TranscribeRequestsCh <- req

	for {
		select {
		case err := <-errorCh:
			return err
		}
	}
}

func (a *TranscribeAction) RunPreActions(request *structs.ActionRequest) error {
	return defaultPreActions(a, request)
}

func (a *TranscribeAction) RunPostActions(request *structs.ActionRequest) error {
	return ErrProceed
}
