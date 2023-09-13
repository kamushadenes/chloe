package latex

import (
	"fmt"

	"github.com/go-latex/latex/drawtex/drawimg"
	"github.com/go-latex/latex/mtex"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *LatexAction) GetNotification() string {
	return fmt.Sprintf("⚗️ Rendering latex: **%s**", a.MustGetParam("formula"))
}

func (a *LatexAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Image)

	dst := drawimg.NewRenderer(obj)
	if err := mtex.Render(dst, a.MustGetParam("formula"), 64, 600, nil); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*response_object_structs.ResponseObject{obj}, nil
}
