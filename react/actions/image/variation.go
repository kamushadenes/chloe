package image

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
)

type VariationAction struct {
	Name   string
	Params map[string]string
}

func (a *VariationAction) GetNotification() string {
	return fmt.Sprintf("🖼️ Generating image: **%s**", a.Params["path"])
}

func (a *VariationAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Image)

	errorCh := make(chan error)

	req := structs.NewVariationRequest()
	req.Context = request.GetContext()
	req.ImagePath = a.Params["path"]
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
