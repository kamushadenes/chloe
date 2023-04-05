package image

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

type ImageAction struct {
	Name   string
	Params string
}

func NewImageAction() structs.Action {
	return &ImageAction{
		Name: "image",
	}
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

func (a *ImageAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Image)

	errorCh := make(chan error)

	req := structs.NewGenerationRequest()
	req.Context = request.GetContext()
	req.Prompt = a.Params
	req.ErrorChannel = errorCh
	req.Count = 1
	// add count
	// req.Count = request.Message.??

	req.Writer = obj

	channels.GenerationRequestsCh <- req

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

func (a *ImageAction) RunPreActions(request *structs.ActionRequest) error {
	return imagePreActions(a, request)
}

func (a *ImageAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
