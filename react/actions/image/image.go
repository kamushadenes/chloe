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
	var objs []*structs.ResponseObject

	for k := 0; k < request.Count; k++ {
		obj := structs.NewResponseObject(structs.Image)

		errorCh := make(chan error)

		req := structs.NewGenerationRequest()
		req.Context = request.GetContext()
		req.Prompt = a.Params
		req.ErrorChannel = errorCh
		req.Count = 1

		req.Writer = obj

		channels.GenerationRequestsCh <- req

		err := <-errorCh

		if err != nil {
			return objs, errors.Wrap(errors.ErrActionFailed, err)
		}

		objs = append(objs, obj)
	}

	return objs, nil
}

func (a *ImageAction) RunPreActions(request *structs.ActionRequest) error {
	return imagePreActions(a, request)
}

func (a *ImageAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
