package image

import (
	"fmt"

	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *VariationAction) GetNotification() string {
	return fmt.Sprintf("🖼️ Generating image: **%s**", a.MustGetParam("path"))
}

func (a *VariationAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	/*
		obj := response_object_structs.NewResponseObject(response_object_structs.Image)

		errorCh := make(chan error)

		req := structs.NewVariationRequest()
		req.Context = request.GetContext()
		req.ImagePath = a.Params["path"]
		req.ErrorChannel = errorCh
		req.Count = request.Count

		req.Writer = obj

		//structs.VariationRequestsCh <- req

		for {
			select {
			case err := <-errorCh:
				if err != nil {
					return nil, errors.Wrap(errors.ErrActionFailed, err)
				}
				return []*structs.ResponseObject{obj}, nil
			}
		}
	*/

	return nil, nil
}
