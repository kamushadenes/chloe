// Code generated by go generate; DO NOT EDIT.
// This file was generated at 2023-09-13 05:26:45.441028 -0300 -03 m=+0.015049167
package scrape

import (
	"fmt"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
)

type ScrapeAction struct {
	Name        string
	Description string
	Params *action_structs.ActionParameterSet
	Extra map[string]interface{}
}
// NewScrapeAction creates a new ScrapeAction with Params initialized
func NewScrapeAction() action_structs.Action {
	var params action_structs.ActionParameterSet
	params.AddParam(&action_structs.ActionParameter{
		Name:        "url",
		Description: "URL to scrape",
		Type:        "string",
		Required:    true,
		Enum:        nil,
	})

	return &ScrapeAction{
		Name:   "scrape",
		Description: "Scrape a website",
		Params: &params,
	}
}

func (a *ScrapeAction) SkipFunctionCall() bool {
	return false
}

// CheckRequiredParams checks if all required params are set
func (a *ScrapeAction) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}
func (a *ScrapeAction) GetName() string {
	return a.Name
}
func (a *ScrapeAction) GetDescription() string {
	return a.Description
}
func (a *ScrapeAction) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
func (a *ScrapeAction) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)
	
	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}	

	return p.Value, err
}

func (a *ScrapeAction) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
func (a *ScrapeAction) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
func (a *ScrapeAction) SetMessage(msg *memory.Message) {}

func (a *ScrapeAction) GetSchema() *functions.FunctionDefinition {
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
