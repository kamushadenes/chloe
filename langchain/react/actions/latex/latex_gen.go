// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-06-17 04:59:23.250627 -0300 -03 m=+0.003651001
package latex

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs"
)

// NewLatexAction creates a new LatexAction with Params initialized
func NewLatexAction() structs.Action {
	return &LatexAction{
		Name:   "latex",
		Params: make(map[string]string),
	}
}

// CheckRequiredParams checks if all required params are set
func (a *LatexAction) CheckRequiredParams() error {
	required := []string{
		"formula",
	}

	for k := range required {
		if _, err := a.GetParam(required[k]); err != nil {
			return errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("required param %s is not set", required[k]))
		}
	}

	return nil
}
func (a *LatexAction) GetName() string {
	return a.Name
}
func (a *LatexAction) SetParam(key, value string) {
	a.Params[key] = value
}
func (a *LatexAction) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
func (a *LatexAction) GetParams() map[string]string {
	return a.Params
}
func (a *LatexAction) SetMessage(msg *memory.Message) {}
func (a *LatexAction) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *LatexAction) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
