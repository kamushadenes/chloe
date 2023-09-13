package image

import (
	"fmt"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/diffusion_models/common"
	base "github.com/kamushadenes/chloe/langchain/diffusion_models/diffusion"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *ImageAction) GetRequiredParams() []string {
	return []string{
		"prompt",
	}
}

func (a *ImageAction) GetNotification() string {
	return fmt.Sprintf("🖼️ Generating image: **%s**", a.MustGetParam("prompt"))
}

func (a *ImageAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	var objs []*response_object_structs.ResponseObject

	dif := base.NewDiffusionWithDefaultModel(config.Diffusion.Provider)

	res, err := dif.GenerateWithContext(request.Context, common.DiffusionMessage{Prompt: a.MustGetParam("prompt")})
	if err != nil {
		return nil, err
	}

	if err != nil {
		return objs, errors.Wrap(errors.ErrActionFailed, err)
	}

	for k := range res.Images {
		ro := response_object_structs.NewResponseObject(response_object_structs.Image)
		ro.Data = res.Images[k]
		ro.HTTPHeader.Add("Content-Type", "image/png")

		objs = append(objs, ro)
	}

	return objs, nil
}

func (a *ImageAction) RunPreActions(request *action_structs.ActionRequest) error {
	return imagePreActions(a, request)
}
