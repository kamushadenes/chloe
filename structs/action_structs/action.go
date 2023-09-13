package action_structs

import (
	"fmt"

	"github.com/kamushadenes/chloe/langchain/actions/functions"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

type Action interface {
	// GetName Get the name of the action
	GetName() string

	// GetDescription Get the description of the action
	GetDescription() string

	// GetNotification  Get the notification message that will be sent when the action is executed
	GetNotification() string

	// SetParam Set the parameters of the action
	SetParam(string, string)

	// GetParam Get the value of a parameter
	GetParam(string) (string, error)

	// GetParams Get all parameters
	GetParams() []*ActionParameter

	// CheckRequiredParams Check if all required parameters are set
	CheckRequiredParams() error

	// SetMessage Set the message that triggered the action
	SetMessage(*memory.Message)

	// Execute The actual implementation of the action
	Execute(*ActionRequest) ([]*response_object_structs.ResponseObject, error)

	// RunPreActions Actions that will be executed before the action is executed
	RunPreActions(*ActionRequest) error

	// RunPostActions Actions that will be executed after the action is executed
	RunPostActions(*ActionRequest) error

	// GetSchema Get the schema of the action to be passed to `functions`
	GetSchema() *functions.FunctionDefinition

	// SkipFunctionCall Check if the action can be automatically executed by the AI
	SkipFunctionCall() bool
}

type ActionParameterSet struct {
	Params []*ActionParameter
}

func (a *ActionParameterSet) AddParam(param *ActionParameter) {
	a.Params = append(a.Params, param)
}

func (a *ActionParameterSet) SetParam(key, value string) {
	for _, p := range a.Params {
		if p.Name == key {
			p.Value = value
			return
		}
	}
}

func (a *ActionParameterSet) GetParams() []*ActionParameter {
	return a.Params
}

func (a *ActionParameterSet) GetParam(name string) (*ActionParameter, error) {
	for _, p := range a.Params {
		if p.Name == name {
			return p, nil
		}
	}

	return nil, fmt.Errorf("param %s not found", name)
}

func (a *ActionParameterSet) CheckRequiredParams() error {
	for _, p := range a.Params {
		if p.Required && p.Value == "" {
			return fmt.Errorf("required param %s is not set", p.Name)
		}
	}
	return nil
}

type ActionParameter struct {
	Name        string   `json:"-"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Enum        []string `json:"enum,omitempty"`
	Required    bool     `json:"-"`
	Value       string   `json:"-"`
}
