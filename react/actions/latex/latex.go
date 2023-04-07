package latex

import (
	"fmt"
	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/mtex"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
)

type LatexAction struct {
	Name   string
	Params map[string]string
}

func (a *LatexAction) GetNotification() string {
	return fmt.Sprintf("⚗️ Rendering latex: **%s**", a.Params["formula"])
}

func (a *LatexAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Image)

	dst := drawimg.NewRenderer(obj)
	if err := mtex.Render(dst, a.Params["formula"], 64, 600, nil); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil
}
