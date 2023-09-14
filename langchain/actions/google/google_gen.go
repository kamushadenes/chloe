// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-14 03:26:37.090361 -0300 -03 m=+0.012401376
package google

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

type GoogleAction struct {
	Name        string
	Description string
	Params      *action_structs.ActionParameterSet
	Extra       map[string]interface{}
}

// NewGoogleAction creates a new GoogleAction with Params initialized
func NewGoogleAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "query",
		Description: "Query to search on Google",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &GoogleAction{
		Name:        "google",
		Description: "Search Google",
		Params:      &params,
	}
}

func (a *GoogleAction) SkipFunctionCall() bool {
	return false
}

// CheckRequiredParams checks if all required params are set
func (a *GoogleAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *GoogleAction) GetName() string {
	return a.Name
}
func (a *GoogleAction) GetDescription() string {
	return a.Description
}
func (a *GoogleAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *GoogleAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)

	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}

	return p.Value, err
}

func (a *GoogleAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *GoogleAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *GoogleAction) SetMessage(msg *memory.Message) {}
func (a *GoogleAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *GoogleAction) GetSchema() *functions.FunctionDefinition {
	params := make(map[string]interface{})

	params["type"] = "object"
	params["required"] = []string{}
	params["properties"] = make(map[string]interface{})

	for k := range a.GetParams() {
		p := a.GetParams()[k]
		params["properties"].(map[string]interface{})[p.Name] = make(map[string]interface{})

		params["properties"].(map[string]interface{})[p.Name].(map[string]interface{})["type"] = p.Type
		params["properties"].(map[string]interface{})[p.Name].(map[string]interface{})["description"] = p.Description
		if p.Enum != nil {
			params["properties"].(map[string]interface{})[p.Name].(map[string]interface{})["enum"] = p.Enum
		}

		if p.Required {
			params["required"] = append(params["required"].([]string), p.Name)
		}
	}

	return &functions.FunctionDefinition{
		Name:        a.GetName(),
		Description: a.GetDescription(),
		Parameters:  params,
	}
}
