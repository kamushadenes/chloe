package file

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	"os"
	"path"
)

type ReadFileAction struct {
	Name   string
	Params map[string]string
}

func (a *ReadFileAction) GetNotification() string {
	return fmt.Sprintf("📄 Writing file: **%s**", a.Params["path"])
}

func (a *ReadFileAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	fname := path.Join(config.React.FileWorkspace, a.Params["path"])

	b, err := os.ReadFile(fname)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Read(b); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil

}
