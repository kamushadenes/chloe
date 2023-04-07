package math

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

type MathAction struct {
	Name   string
	Params map[string]string
}

func (a *MathAction) GetNotification() string {
	return fmt.Sprintf("🧮 Executing calculation: **%s**", strings.ReplaceAll(a.Params["expression"], "*", "\\*"))
}

func (a *MathAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	expr := strings.ReplaceAll(a.Params["expression"], ",", "")

	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	result, err := expression.Evaluate(make(map[string]interface{}))
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Write([]byte(fmt.Sprintf("%v", result))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil
}
