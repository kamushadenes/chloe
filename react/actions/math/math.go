package math

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"strings"
)

type CalculateAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewCalculateAction() structs.Action {
	return &CalculateAction{
		Name: "calculate",
	}
}

func (a *CalculateAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *CalculateAction) GetWriters() []io.WriteCloser {
	return a.Writers
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

func (a *CalculateAction) Execute(request *structs.ActionRequest) error {
	expr := strings.ReplaceAll(a.Params, ",", "")

	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return errors.Wrap(errors.ErrActionFailed, err)
	}

	result, err := expression.Evaluate(make(map[string]interface{}))
	if err != nil {
		return errors.Wrap(errors.ErrActionFailed, err)
	}

	for _, w := range a.Writers {
		_, err := w.Write([]byte(fmt.Sprintf("%v", result)))
		if err != nil {
			return errors.Wrap(errors.ErrActionFailed, err)
		}
	}

	return nil
}

func (a *CalculateAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *CalculateAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
