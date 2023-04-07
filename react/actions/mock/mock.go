package mock

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
)

type MockAction struct {
	Name   string
	Params map[string]string
}

func (a *MockAction) GetNotification() string {
	return fmt.Sprintf("❓ Mock Action: **%s**", a.Params["foo"])
}

func (a *MockAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	if request.Params["foo"] == "err" {
		return nil, errors.ErrMock
	}

	if _, err := obj.Write([]byte(a.Params["foo"])); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil
}

func (a *MockAction) RunPreActions(request *structs.ActionRequest) error {
	if request.Params["foo"] == "errPre" {
		return errors.ErrMock
	}
	return nil
}

func (a *MockAction) RunPostActions(request *structs.ActionRequest) error {
	if request.Params["foo"] == "errPost" {
		return errors.ErrMock
	}
	return nil
}
