// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-04-07 00:50:33.614505 -0300 -03 m=+0.003125334
package transcribe

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"

	"github.com/kamushadenes/chloe/structs"
)
// NewTranscribeAction creates a new TranscribeAction with Params initialized
func NewTranscribeAction() structs.Action {
	return &TranscribeAction{
		Name:   "transcribe",
		Params: make(map[string]string),
	}
}

// CheckRequiredParams checks if all required params are set
func (a *TranscribeAction) CheckRequiredParams() error {
	required := []string{
		"path",
	}

	for k := range required {
		if _, err := a.GetParam(required[k]); err != nil {
			return errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("required param %s is not set", required[k]))
		}
	}

	return nil
}
func (a *TranscribeAction) GetName() string {
	return a.Name
}
func (a *TranscribeAction) SetParam(key, value string) {
	a.Params[key] = value
}
func (a *TranscribeAction) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
func (a *TranscribeAction) GetParams() map[string]string {
	return a.Params
}
func (a *TranscribeAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}