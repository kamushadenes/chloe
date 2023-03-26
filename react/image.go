package react

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type ImageAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewImageAction() Action {
	return &ImageAction{
		Name: "image",
	}
}

func (a *ImageAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *ImageAction) GetWriters() []io.WriteCloser {
	return a.Writers
}
func (a *ImageAction) GetName() string {
	return a.Name
}

func (a *ImageAction) GetNotification() string {
	return fmt.Sprintf("🖼️ Generating image: **%s**", a.Params)
}

func (a *ImageAction) SetParams(params string) {
	a.Params = params
}

func (a *ImageAction) GetParams() string {
	return a.Params
}

func (a *ImageAction) SetMessage(message *memory.Message) {}

func (a *ImageAction) Execute(request *structs.ActionRequest) error {
	errorCh := make(chan error)

	req := structs.NewGenerationRequest()
	req.Context = request.GetContext()
	req.Prompt = a.Params
	req.ErrorChannel = errorCh

	req.Writers = a.Writers

	channels.GenerationRequestsCh <- req

	for {
		select {
		case err := <-errorCh:
			return err
		}
	}
}

func (a *ImageAction) RunPreActions(request *structs.ActionRequest) error {
	return imagePreActions(a, request)
}

func (a *ImageAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
