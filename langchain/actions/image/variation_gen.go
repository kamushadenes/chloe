// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-21 13:57:46.184236 -0300 -03 m=+0.010620626
package image

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

type VariationAction struct {
	Name        string
	Description string
	Params      *action_structs.ActionParameterSet
	Extra       map[string]interface{}
}

// NewVariationAction creates a new VariationAction with Params initialized
func NewVariationAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "path",
		Description: "Path to the image to vary",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &VariationAction{
		Name:        "variation",
		Description: "Generate a variation of an image",
		Params:      &params,
	}
}

func (a *VariationAction) SkipFunctionCall() bool {
	return true
}

// CheckRequiredParams checks if all required params are set
func (a *VariationAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *VariationAction) GetName() string {
	return a.Name
}
func (a *VariationAction) GetDescription() string {
	return a.Description
}
func (a *VariationAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *VariationAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)

	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}

	return p.Value, err
}

func (a *VariationAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *VariationAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *VariationAction) SetMessage(msg *memory.Message) {}
func (a *VariationAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *VariationAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *VariationAction) GetSchema() *functions.FunctionDefinition {
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
