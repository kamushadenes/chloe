// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-14 02:07:33.069424 -0300 -03 m=+0.009239626
package mock

import (
	"fmt"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
)

type MockAction struct {
	Name        string
	Description string
	Params *action_structs.ActionParameterSet
	Extra map[string]interface{}
}
// NewMockAction creates a new MockAction with Params initialized
func NewMockAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "foo",
		Description: "Foo action_structs.ActionParametereter",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &MockAction{
		Name:   "mock",
		Description: "Mock a message",
		Params: &params,
	}
}

func (a *MockAction) SkipFunctionCall() bool {
	return true
}

// CheckRequiredParams checks if all required params are set
func (a *MockAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *MockAction) GetName() string {
	return a.Name
}
func (a *MockAction) GetDescription() string {
	return a.Description
}
func (a *MockAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *MockAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)
	
	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}	

	return p.Value, err
}

func (a *MockAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *MockAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *MockAction) SetMessage(msg *memory.Message) {}

func (a *MockAction) GetSchema() *functions.FunctionDefinition {
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
