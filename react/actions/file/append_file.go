package file

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	"os"
	"path"
)

type AppendFileAction struct {
	Name   string
	Params map[string]string
}

func (a *AppendFileAction) GetNotification() string {
	return fmt.Sprintf("📄 Appending file: **%s**", a.Params["path"])
}

func (a *AppendFileAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	fname := path.Join(config.React.FileWorkspace, a.Params["path"])
	content := a.Params["content"]

	if err := os.MkdirAll(path.Dir(fname), 0755); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}
	defer f.Close()

	if _, err := f.WriteString(content); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Write([]byte(fmt.Sprintf("File appended: **%s**", fname))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil

}
