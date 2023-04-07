package image

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
)

type ImageAction struct {
	Name   string
	Params map[string]string
}

func (a *ImageAction) GetRequiredParams() []string {
	return []string{
		"prompt",
	}
}

func (a *ImageAction) GetNotification() string {
	return fmt.Sprintf("🖼️ Generating image: **%s**", a.Params["prompt"])
}

func (a *ImageAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	var objs []*structs.ResponseObject

	writer := structs.NewMockWriter()

	errorCh := make(chan error)

	req := structs.NewGenerationRequest()
	req.Context = request.GetContext()
	req.Prompt = a.Params["prompt"]
	req.ErrorChannel = errorCh
	req.Count = request.Count

	req.Writer = writer

	channels.GenerationRequestsCh <- req

	err := <-errorCh

	if err != nil {
		return objs, errors.Wrap(errors.ErrActionFailed, err)
	}

	objs = append(objs, writer.GetObjects()...)

	return objs, nil
}

func (a *ImageAction) RunPreActions(request *structs.ActionRequest) error {
	return imagePreActions(a, request)
}
