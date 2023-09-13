package math

import (
	"fmt"
	"strings"

	"github.com/antonmedv/expr"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *MathAction) GetNotification() string {
	return fmt.Sprintf("🧮 Executing calculation: **%s**", strings.ReplaceAll(a.MustGetParam("expression"), "*", "\\*"))
}

func (a *MathAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	expression := strings.ReplaceAll(a.MustGetParam("expression"), ",", "")

	program, err := expr.Compile(expression, expr.Env(nil))
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	result, err := expr.Run(program, nil)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Write([]byte(fmt.Sprintf("%v", result))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*response_object_structs.ResponseObject{obj}, nil
}
