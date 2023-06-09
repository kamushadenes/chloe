// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-04-07 00:50:33.614233 -0300 -03 m=+0.002852543
package scrape

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
)
// NewScrapeAction creates a new ScrapeAction with Params initialized
func NewScrapeAction() structs.Action {
	return &ScrapeAction{
		Name:   "scrape",
		Params: make(map[string]string),
	}
}

// CheckRequiredParams checks if all required params are set
func (a *ScrapeAction) CheckRequiredParams() error {
	required := []string{
		"url",
	}

	for k := range required {
		if _, err := a.GetParam(required[k]); err != nil {
			return errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("required param %s is not set", required[k]))
		}
	}

	return nil
}
func (a *ScrapeAction) GetName() string {
	return a.Name
}
func (a *ScrapeAction) SetParam(key, value string) {
	a.Params[key] = value
}
func (a *ScrapeAction) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
func (a *ScrapeAction) GetParams() map[string]string {
	return a.Params
}
func (a *ScrapeAction) SetMessage(msg *memory.Message) {}
