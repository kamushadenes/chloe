// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-04-07 00:50:33.613028 -0300 -03 m=+0.001648334
package google

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)
// NewGoogleAction creates a new GoogleAction with Params initialized
func NewGoogleAction() structs.Action {
	return &GoogleAction{
		Name:   "google",
		Params: make(map[string]string),
	}
}

// CheckRequiredParams checks if all required params are set
func (a *GoogleAction) CheckRequiredParams() error {
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
func (a *GoogleAction) GetName() string {
	return a.Name
}
func (a *GoogleAction) SetParam(key, value string) {
	a.Params[key] = value
}
func (a *GoogleAction) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
func (a *GoogleAction) GetParams() map[string]string {
	return a.Params
}
func (a *GoogleAction) SetMessage(msg *memory.Message) {}
func (a *GoogleAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
