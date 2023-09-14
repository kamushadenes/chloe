package file

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

func (a *AppendFileAction) GetNotification() string {
	return fmt.Sprintf("📄 Appending file: **%s**", a.MustGetParam("path"))
}

func (a *AppendFileAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	fname := filepath.Join(config.React.FileWorkspace, a.MustGetParam("path"))
	content := a.MustGetParam("content")

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

	if _, err := obj.Write(action_structs.SuccessMessage); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*response_object_structs.ResponseObject{obj}, nil
}
