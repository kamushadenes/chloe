// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-04-07 00:50:33.613151 -0300 -03 m=+0.001770751
package image

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)
// NewImageAction creates a new ImageAction with Params initialized
func NewImageAction() structs.Action {
	return &ImageAction{
		Name:   "image",
		Params: make(map[string]string),
	}
}

// CheckRequiredParams checks if all required params are set
func (a *ImageAction) CheckRequiredParams() error {
	required := []string{
		"prompt",
	}

	for k := range required {
		if _, err := a.GetParam(required[k]); err != nil {
			return errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("required param %s is not set", required[k]))
		}
	}

	return nil
}
func (a *ImageAction) GetName() string {
	return a.Name
}
func (a *ImageAction) SetParam(key, value string) {
	a.Params[key] = value
}
func (a *ImageAction) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
func (a *ImageAction) GetParams() map[string]string {
	return a.Params
}
func (a *ImageAction) SetMessage(msg *memory.Message) {}
func (a *ImageAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
