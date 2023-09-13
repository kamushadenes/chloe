// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-12 20:56:33.555678 -0300 -03 m=+0.007733710
package file

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
)

type DeleteFileAction struct {
	Name        string
	Description string
	Params *action_structs.ActionParameterSet
	Extra map[string]interface{}
}
// NewDeleteFileAction creates a new DeleteFileAction with Params initialized
func NewDeleteFileAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "path",
		Description: "Path to the file to delete",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &DeleteFileAction{
		Name:   "delete_file",
		Description: "Delete a file",
		Params: &params,
	}
}

// CheckRequiredParams checks if all required params are set
func (a *DeleteFileAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *DeleteFileAction) GetName() string {
	return a.Name
}
func (a *DeleteFileAction) GetDescription() string {
	return a.Description
}
func (a *DeleteFileAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *DeleteFileAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)
	
	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}	

	return p.Value, err
}

func (a *DeleteFileAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *DeleteFileAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *DeleteFileAction) SetMessage(msg *memory.Message) {}
func (a *DeleteFileAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *DeleteFileAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *DeleteFileAction) GetSchema() *functions.FunctionDefinition {
	params := make(map[string]interface{})
	
	params["parameters"] = make(map[string]interface{})
	
	params["parameters"].(map[string]interface{})["type"] = "object"
	params["parameters"].(map[string]interface{})["required"] = []string{}
	
	for k := range a.GetParams() {
		p := a.GetParams()[k]
		params["parameters"].(map[string]interface{})[p.Name] = make(map[string]interface{})	
		
		params["parameters"].(map[string]interface{})[p.Name].(map[string]interface{})["type"] = p.Type
		params["parameters"].(map[string]interface{})[p.Name].(map[string]interface{})["description"] = p.Description
		params["parameters"].(map[string]interface{})[p.Name].(map[string]interface{})["enum"] = p.Enum
		
		if p.Required {
			params["parameters"].(map[string]interface{})["required"] = append(params["parameters"].(map[string]interface{})["required"].([]string), p.Name)
		}
	}

	return &functions.FunctionDefinition{
		Name:        a.GetName(),
		Description: a.GetDescription(),
		Parameters:      a.GetParams(),
	}
}
