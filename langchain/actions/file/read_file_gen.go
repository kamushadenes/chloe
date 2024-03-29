// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-21 13:57:46.182846 -0300 -03 m=+0.009230168
package file

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

type ReadFileAction struct {
	Name        string
	Description string
	Params      *action_structs.ActionParameterSet
	Extra       map[string]interface{}
}

// NewReadFileAction creates a new ReadFileAction with Params initialized
func NewReadFileAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "path",
		Description: "Path to the file to read",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &ReadFileAction{
		Name:        "read_file",
		Description: "Read a file",
		Params:      &params,
	}
}

func (a *ReadFileAction) SkipFunctionCall() bool {
	return false
}

// CheckRequiredParams checks if all required params are set
func (a *ReadFileAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *ReadFileAction) GetName() string {
	return a.Name
}
func (a *ReadFileAction) GetDescription() string {
	return a.Description
}
func (a *ReadFileAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *ReadFileAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)

	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}

	return p.Value, err
}

func (a *ReadFileAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *ReadFileAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *ReadFileAction) SetMessage(msg *memory.Message) {}
func (a *ReadFileAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *ReadFileAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *ReadFileAction) GetSchema() *functions.FunctionDefinition {
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
