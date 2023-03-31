package mock

import (
	"fmt"
	"github.com/kamushadenes/chloe/memory"
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/react/errors"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type MockAction struct {
	Name    string
	Params  string
	Writers []io.WriteCloser
}

func NewMockAction() structs2.Action {
	return &MockAction{
		Name: "mock",
	}
}

func (a *MockAction) SetWriters(writers []io.WriteCloser) {
	a.Writers = writers
}

func (a *MockAction) GetWriters() []io.WriteCloser {
	return a.Writers
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

func (a *MockAction) Execute(request *structs.ActionRequest) error {
	if request.Params == "err" {
		return errors.ErrMock
	}

	for _, w := range a.Writers {
		_, err := w.Write([]byte(request.Params))
		if err != nil {
			return err
		}
	}
	return nil
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
