// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-21 13:57:46.18165 -0300 -03 m=+0.008034709
package file

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

type AppendFileAction struct {
	Name        string
	Description string
	Params      *action_structs.ActionParameterSet
	Extra       map[string]interface{}
}

// NewAppendFileAction creates a new AppendFileAction with Params initialized
func NewAppendFileAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "path",
		Description: "Path to the file to append to",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})
	params.AddParam(&action_structs.ActionParameter{
		Name:        "content",
		Description: "Content to append to the file",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &AppendFileAction{
		Name:        "append_file",
		Description: "Append content to a file",
		Params:      &params,
	}
}

func (a *AppendFileAction) SkipFunctionCall() bool {
	return false
}

// CheckRequiredParams checks if all required params are set
func (a *AppendFileAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *AppendFileAction) GetName() string {
	return a.Name
}
func (a *AppendFileAction) GetDescription() string {
	return a.Description
}
func (a *AppendFileAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *AppendFileAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)

	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}

	return p.Value, err
}

func (a *AppendFileAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *AppendFileAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *AppendFileAction) SetMessage(msg *memory.Message) {}
func (a *AppendFileAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *AppendFileAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *AppendFileAction) GetSchema() *functions.FunctionDefinition {
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
