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

func (a *ReadFileAction) GetNotification() string {
	return fmt.Sprintf("📄 Writing file: **%s**", a.MustGetParam("path"))
}

func (a *ReadFileAction) Execute(request *action_structs.ActionRequest) ([]*response_object_structs.ResponseObject, error) {
	obj := response_object_structs.NewResponseObject(response_object_structs.Text)

	fname := path.Join(config.React.FileWorkspace, a.MustGetParam("path"))

	b, err := os.ReadFile(fname)
	if err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	if _, err := obj.Read(b); err != nil {
		return nil, errors.Wrap(errors.ErrActionFailed, err)
	}

	return []*response_object_structs.ResponseObject{obj}, nil

}
