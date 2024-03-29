// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-21 13:57:46.188004 -0300 -03 m=+0.014388626
package tts

import (
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

type TTSAction struct {
	Name        string
	Description string
	Params      *action_structs.ActionParameterSet
	Extra       map[string]interface{}
}

// NewTTSAction creates a new TTSAction with Params initialized
func NewTTSAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "text",
		Description: "Text to convert to speech",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &TTSAction{
		Name:        "tts",
		Description: "Text to speech",
		Params:      &params,
	}
}

func (a *TTSAction) SkipFunctionCall() bool {
	return true
}

// CheckRequiredParams checks if all required params are set
func (a *TTSAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *TTSAction) GetName() string {
	return a.Name
}
func (a *TTSAction) GetDescription() string {
	return a.Description
}
func (a *TTSAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *TTSAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)

	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}

	return p.Value, err
}

func (a *TTSAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *TTSAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *TTSAction) SetMessage(msg *memory.Message) {}
func (a *TTSAction) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
func (a *TTSAction) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}

func (a *TTSAction) GetSchema() *functions.FunctionDefinition {
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
