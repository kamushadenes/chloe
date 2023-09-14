// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-14 02:07:33.069598 -0300 -03 m=+0.009413584
package news

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
)

type NewsAction struct {
	Name        string
	Description string
	Params *action_structs.ActionParameterSet
	Extra map[string]interface{}
}
// NewNewsAction creates a new NewsAction with Params initialized
func NewNewsAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "query",
		Description: "Query to search news",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &NewsAction{
		Name:   "news",
		Description: "Get news",
		Params: &params,
	}
}

func (a *NewsAction) SkipFunctionCall() bool {
	return false
}

// CheckRequiredParams checks if all required params are set
func (a *NewsAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *NewsAction) GetName() string {
	return a.Name
}
func (a *NewsAction) GetDescription() string {
	return a.Description
}
func (a *NewsAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *NewsAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)
	
	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}	

	return p.Value, err
}

func (a *NewsAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *NewsAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *NewsAction) SetMessage(msg *memory.Message) {}
func (a *NewsAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *NewsAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *NewsAction) GetSchema() *functions.FunctionDefinition {
	params:= make(map[string]interface{})
	
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
