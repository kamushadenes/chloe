package file

import (
	"fmt"
	"os"
	"path"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *WriteFileAction) GetNotification() string {
	return fmt.Sprintf("📄 Writing file: **%s**", a.MustGetParam("path"))
}

func (a *WriteFileAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	fname := path.Join(config.React.FileWorkspace, a.MustGetParam("path"))
	content := a.MustGetParam("content")

	if err := os.MkdirAll(path.Dir(fname), 0755); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if err := os.WriteFile(fname, []byte(content), 0644); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Write([]byte(fmt.Sprintf("File written: **%s**", fname))); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*response_object_structs.ResponseObject{obj}, nil

}
