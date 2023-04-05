package mock

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)

type MockAction struct {
	Name   string
	Params string
}

func NewMockAction() structs.Action {
	return &MockAction{
		Name: "mock",
	}
}

func (a *MockAction) GetName() string {
	return a.Name
}

func (a *MockAction) GetNotification() string {
	return fmt.Sprintf("❓ Mock Action: **%s**", a.Params)
}

func (a *MockAction) SetParams(params string) {
	a.Params = params
}

func (a *MockAction) GetParams() string {
	return a.Params
}

func (a *MockAction) SetMessage(message *memory.Message) {}

func (a *MockAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	if request.Params == "err" {
		return nil, errors.ErrMock
	}

	if _, err := obj.Write([]byte(request.Params)); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil
}

func (a *MockAction) RunPreActions(request *structs.ActionRequest) error {
	if request.Params == "errPre" {
		return errors.ErrMock
	}
	return nil
}

func (a *MockAction) RunPostActions(request *structs.ActionRequest) error {
	if request.Params == "errPost" {
		return errors.ErrMock
	}
	return nil
}
