// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-04-07 00:50:33.613997 -0300 -03 m=+0.002616709
package news

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)
// NewNewsAction creates a new NewsAction with Params initialized
func NewNewsAction() structs.Action {
	return &NewsAction{
		Name:   "news",
		Params: make(map[string]string),
	}
}

// CheckRequiredParams checks if all required params are set
func (a *NewsAction) CheckRequiredParams() error {
	required := []string{
		"query",
	}

	for k := range required {
		if _, err := a.GetParam(required[k]); err != nil {
			return errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("required param %s is not set", required[k]))
		}
	}

	return nil
}
func (a *NewsAction) GetName() string {
	return a.Name
}
func (a *NewsAction) SetParam(key, value string) {
	a.Params[key] = value
}
func (a *NewsAction) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
func (a *NewsAction) GetParams() map[string]string {
	return a.Params
}
func (a *NewsAction) SetMessage(msg *memory.Message) {}
func (a *NewsAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *NewsAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
