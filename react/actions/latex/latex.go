package latex

import (
	"fmt"
	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/mtex"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

type LatexAction struct {
	Name   string
	Params string
}

func NewLatexAction() structs.Action {
	return &LatexAction{
		Name: "latex",
	}
}

func (a *LatexAction) GetName() string {
	return a.Name
}

func (a *LatexAction) GetNotification() string {
	return fmt.Sprintf("⚗️ Rendering latex: **%s**", a.Params)
}

func (a *LatexAction) SetParams(params string) {
	a.Params = params
}

func (a *LatexAction) GetParams() string {
	return a.Params
}

func (a *LatexAction) SetMessage(message *memory.Message) {}

func (a *LatexAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Image)

	dst := drawimg.NewRenderer(obj)
	if err := mtex.Render(dst, request.Params, 64, 600, nil); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil
}

func (a *LatexAction) RunPreActions(request *structs.ActionRequest) error {
	return latexPreActions(a, request)
}

func (a *LatexAction) RunPostActions(request *structs.ActionRequest) error {
	return nil
}
