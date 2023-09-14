// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-14 03:26:37.090123 -0300 -03 m=+0.012163417
package transcribe

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
)

type TranscribeAction struct {
	Name        string
	Description string
	Params *action_structs.ActionParameterSet
	Extra map[string]interface{}
}
// NewTranscribeAction creates a new TranscribeAction with Params initialized
func NewTranscribeAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "path",
		Description: "Path to the audio file to transcribe",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &TranscribeAction{
		Name:   "transcribe",
		Description: "Transcribe an audio file",
		Params: &params,
	}
}

func (a *TranscribeAction) SkipFunctionCall() bool {
	return true
}

// CheckRequiredParams checks if all required params are set
func (a *TranscribeAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *TranscribeAction) GetName() string {
	return a.Name
}
func (a *TranscribeAction) GetDescription() string {
	return a.Description
}
func (a *TranscribeAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *TranscribeAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)
	
	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}	

	return p.Value, err
}

func (a *TranscribeAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *TranscribeAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *TranscribeAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *TranscribeAction) GetSchema() *functions.FunctionDefinition {
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
