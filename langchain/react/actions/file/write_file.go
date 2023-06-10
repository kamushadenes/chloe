package file

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	"os"
	"path"
)

type WriteFileAction struct {
	Name   string
	Params map[string]string
}

func (a *WriteFileAction) GetNotification() string {
	return fmt.Sprintf("📄 Writing file: **%s**", a.Params["path"])
}

func (a *WriteFileAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	fname := path.Join(config.React.FileWorkspace, a.Params["path"])
	content := a.Params["content"]

	if err := os.MkdirAll(path.Dir(fname), 0755); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if err := os.WriteFile(fname, []byte(content), 0644); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Write([]byte(fmt.Sprintf("File written: **%s**", fname))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil

}
