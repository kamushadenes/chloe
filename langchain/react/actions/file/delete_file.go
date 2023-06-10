package file

import (
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	"os"
	"path"
)

type DeleteFileAction struct {
	Name   string
	Params map[string]string
}

func (a *DeleteFileAction) GetNotification() string {
	return fmt.Sprintf("📄 Deleting file: **%s**", a.Params["path"])
}

func (a *DeleteFileAction) Execute(request *structs.ActionRequest) ([]*structs.ResponseObject, error) {
	obj := structs.NewResponseObject(structs.Text)

	fname := path.Join(config.React.FileWorkspace, a.Params["path"])

	if err := os.Remove(fname); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Write([]byte(fmt.Sprintf("File deleted: **%s**", fname))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*structs.ResponseObject{obj}, nil

}
