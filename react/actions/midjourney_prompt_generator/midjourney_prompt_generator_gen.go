// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-04-07 00:50:33.61376 -0300 -03 m=+0.002379793
package midjourney_prompt_generator

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)
// NewMidjourneyPromptGeneratorAction creates a new MidjourneyPromptGeneratorAction with Params initialized
func NewMidjourneyPromptGeneratorAction() structs.Action {
	return &MidjourneyPromptGeneratorAction{
		Name:   "midjourney_prompt_generator",
		Params: make(map[string]string),
	}
}

// CheckRequiredParams checks if all required params are set
func (a *MidjourneyPromptGeneratorAction) CheckRequiredParams() error {
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
func (a *MidjourneyPromptGeneratorAction) GetName() string {
	return a.Name
}
func (a *MidjourneyPromptGeneratorAction) SetParam(key, value string) {
	a.Params[key] = value
}
func (a *MidjourneyPromptGeneratorAction) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
func (a *MidjourneyPromptGeneratorAction) GetParams() map[string]string {
	return a.Params
}
func (a *MidjourneyPromptGeneratorAction) SetMessage(msg *memory.Message) {}
func (a *MidjourneyPromptGeneratorAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *MidjourneyPromptGeneratorAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}