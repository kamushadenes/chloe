package mock

import (
	"fmt"

	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *MockAction) GetNotification() string {
	return fmt.Sprintf("❓ Mock Action: **%s**", a.MustGetParam("foo"))
}

func (a *MockAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	if request.Params["foo"] == "err" {
		return nil, errors.ErrMock
	}

	if _, err := obj.Write([]byte(a.MustGetParam("foo"))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*response_object_structs.ResponseObject{obj}, nil
}

func (a *MockAction) RunPreActions(request *action_structs.ActionRequest) error {
	if request.Params["foo"] == "errPre" {
		return errors.ErrMock
	}
	return nil
}

func (a *MockAction) RunPostActions(request *action_structs.ActionRequest) error {
	if request.Params["foo"] == "errPost" {
		return errors.ErrMock
	}
	return nil
}
