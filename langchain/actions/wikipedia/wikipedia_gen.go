// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-13 05:26:45.443195 -0300 -03 m=+0.017216542
package wikipedia

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
)

type WikipediaAction struct {
	Name        string
	Description string
	Params *action_structs.ActionParameterSet
	Extra map[string]interface{}
}
// NewWikipediaAction creates a new WikipediaAction with Params initialized
func NewWikipediaAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "query",
		Description: "Query to search on Wikipedia",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &WikipediaAction{
		Name:   "wikipedia",
		Description: "Search Wikipedia",
		Params: &params,
	}
}

func (a *WikipediaAction) SkipFunctionCall() bool {
	return false
}

// CheckRequiredParams checks if all required params are set
func (a *WikipediaAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *WikipediaAction) GetName() string {
	return a.Name
}
func (a *WikipediaAction) GetDescription() string {
	return a.Description
}
func (a *WikipediaAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *WikipediaAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)
	
	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}	

	return p.Value, err
}

func (a *WikipediaAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *WikipediaAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *WikipediaAction) SetMessage(msg *memory.Message) {}
func (a *WikipediaAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *WikipediaAction) GetSchema() *functions.FunctionDefinition {
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