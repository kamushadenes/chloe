package react

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type VariationAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewVariationAction() *VariationAction {
	return &VariationAction{
		Name: "image",
	}
}

func (a *VariationAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *VariationAction) GetWriters() []io.WriteCloser {
	return a.Writers
}

func (a *VariationAction) GetName() string {
	return a.Name
}

func (a *VariationAction) GetNotification() string {
	return fmt.Sprintf("🖼️ Generating image: **%s**", a.Params)
}

func (a *VariationAction) SetParams(params string) {
	a.Params = params
}

func (a *VariationAction) GetParams() string {
	return a.Params
}

func (a *VariationAction) SetUser(user *memory.User)          {}
func (a *VariationAction) SetMessage(message *memory.Message) {}

func (a *VariationAction) Execute(request *structs.ActionRequest) error {
	errorCh := make(chan error)

	req := structs.NewVariationRequest()
	req.Context = request.GetContext()
	req.ImagePath = a.Params
	req.ErrorChannel = errorCh

	req.Writers = a.Writers

	channels.VariationRequestsCh <- req

	for {
		select {
		case err := <-errorCh:
			return err
		}
	}
}

func (a *VariationAction) RunPreActions(request *structs.ActionRequest) error {
	return imagePreActions(a, request)
}

func (a *VariationAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
