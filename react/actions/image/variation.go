package image

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

type VariationAction struct {
	Name   string
	Params string
}

func NewVariationAction() structs.Action {
	return &VariationAction{
		Name: "image",
	}
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

func (a *VariationAction) SetMessage(message *memory.Message) {}

func (a *VariationAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Image)

	errorCh := make(chan error)

	req := structs.NewVariationRequest()
	req.Context = request.GetContext()
	req.ImagePath = a.Params
	req.ErrorChannel = errorCh
	req.Count = request.Count

	req.Writer = obj

	channels.VariationRequestsCh <- req

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

func (a *VariationAction) RunPreActions(request *structs.ActionRequest) error {
	return imagePreActions(a, request)
}

func (a *VariationAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
