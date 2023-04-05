package math

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

type CalculateAction struct {
	Name   string
	Params string
}

func NewCalculateAction() structs.Action {
	return &CalculateAction{
		Name: "calculate",
	}
}

func (a *CalculateAction) GetName() string {
	return a.Name
}

func (a *CalculateAction) GetNotification() string {
	return fmt.Sprintf("🧮 Executing calculation: **%s**", strings.ReplaceAll(a.Params, "*", "\\*"))
}

func (a *CalculateAction) SetParams(params string) {
	a.Params = params
}

func (a *CalculateAction) GetParams() string {
	return a.Params
}

func (a *CalculateAction) SetMessage(message *memory.Message) {}

func (a *CalculateAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	expr := strings.ReplaceAll(a.Params, ",", "")

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

func (a *CalculateAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *CalculateAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
