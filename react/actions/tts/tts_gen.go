// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-04-07 00:50:33.614638 -0300 -03 m=+0.003257626
package tts

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)
// NewTTSAction creates a new TTSAction with Params initialized
func NewTTSAction() structs.Action {
	return &TTSAction{
		Name:   "tts",
		Params: make(map[string]string),
	}
}

// CheckRequiredParams checks if all required params are set
func (a *TTSAction) CheckRequiredParams() error {
	required := []string{
		"text",
	}

	for k := range required {
		if _, err := a.GetParam(required[k]); err != nil {
			return errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("required param %s is not set", required[k]))
		}
	}

	return nil
}
func (a *TTSAction) GetName() string {
	return a.Name
}
func (a *TTSAction) SetParam(key, value string) {
	a.Params[key] = value
}
func (a *TTSAction) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
func (a *TTSAction) GetParams() map[string]string {
	return a.Params
}
func (a *TTSAction) SetMessage(msg *memory.Message) {}
func (a *TTSAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *TTSAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
