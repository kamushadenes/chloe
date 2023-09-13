//go:build ignore

package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/kamushadenes/chloe/structs/action_structs"
)

type generatedAction struct {
	Name               string
	Description        string
	OverrideStructName string
	Package            string
	Params             []*action_structs.ActionParameter
	SkipNewAction      bool
	SkipGetName        bool
	SkipGetDescription bool
	SkipParams         bool
	SkipSetParam       bool
	SkipGetParam       bool
	SkipGetParams      bool
	SkipRunPreActions  bool
	SkipRunPostActions bool
	SkipSetMessage     bool
}

var actionsToGenerate = []*generatedAction{
	{
		Name:        "append_file",
		Description: "Append content to a file",
		Package:     "file",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "path",
				Description: "Path to the file to append to",
				Type:        "string",
				Required:    true,
			},
			{
				Name:        "content",
				Description: "Content to append to the file",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "delete_file",
		Description: "Delete a file",
		Package:     "file",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "path",
				Description: "Path to the file to delete",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "read_file",
		Description: "Read a file",
		Package:     "file",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "path",
				Description: "Path to the file to read",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "write_file",
		Description: "Write content to a file",
		Package:     "file",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "path",
				Description: "Path to the file to write to",
				Type:        "string",
				Required:    true,
			},
			{
				Name:        "content",
				Description: "Content to write to the file",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "image",
		Description: "Generate an image from a prompt",
		Package:     "image",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "prompt",
				Description: "Prompt to generate",
				Type:        "string",
				Required:    true,
			},
		},
		SkipRunPreActions: true,
	},
	{
		Name:        "variation",
		Description: "Generate a variation of an image",
		Package:     "image",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "path",
				Description: "Path to the image to vary",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "latex",
		Description: "Render a LaTeX formula",
		Package:     "latex",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "formula",
				Description: "LaTeX formula to render",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "math",
		Description: "Evaluate a mathematical expression",
		Package:     "math",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "expression",
				Description: "Mathematical expression to evaluate",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "mock",
		Description: "Mock a message",
		Package:     "mock",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "foo",
				Description: "Foo action_structs.ActionParametereter",
				Type:        "string",
				Required:    true,
			},
		},
		SkipRunPreActions:  true,
		SkipRunPostActions: true,
	},
	{
		Name:        "news",
		Description: "Get news",
		Package:     "news",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "query",
				Description: "Query to search news",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "news_by_country",
		Description: "Get news by country",
		Package:     "news",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "country",
				Description: "Country to get news from",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "scrape",
		Description: "Scrape a website",
		Package:     "scrape",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "url",
				Description: "URL to scrape",
				Type:        "string",
				Required:    true,
			},
		},
		SkipRunPreActions:  true,
		SkipRunPostActions: true,
	},
	{
		Name:        "transcribe",
		Description: "Transcribe an audio file",
		Package:     "transcribe",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "path",
				Description: "Path to the audio file to transcribe",
				Type:        "string",
				Required:    true,
			},
		},
		SkipRunPostActions: true,
		SkipSetMessage:     true,
	},
	{
		Name:        "google",
		Description: "Search Google",
		Package:     "google",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "query",
				Description: "Query to search on Google",
				Type:        "string",
				Required:    true,
			},
		},
		SkipRunPostActions: true,
	},
	{
		Name:               "tts",
		Description:        "Text to speech",
		Package:            "tts",
		OverrideStructName: "TTS",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "text",
				Description: "Text to convert to speech",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "wikipedia",
		Description: "Search Wikipedia",
		Package:     "wikipedia",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "query",
				Description: "Query to search on Wikipedia",
				Type:        "string",
				Required:    true,
			},
		},
		SkipRunPostActions: true,
	},
	{
		Name:        "youtube_summarize",
		Description: "Summarize a YouTube video",
		Package:     "youtube",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "url",
				Description: "URL to summarize",
				Type:        "string",
				Required:    true,
			},
		},
	},
	{
		Name:        "youtube_transcribe",
		Description: "Transcribe a YouTube video",
		Package:     "youtube",
		Params: []*action_structs.ActionParameter{
			{
				Name:        "url",
				Description: "URL to transcribe",
				Type:        "string",
				Required:    true,
			},
		},
	},
}

func main() {
	for k := range actionsToGenerate {
		act := actionsToGenerate[k]

		fmt.Printf("Generating methods for action `%s` in package `%s` ...\n", act.Name, act.Package)

		f, err := os.OpenFile(fmt.Sprintf("%s/%s_gen.go", act.Package, act.Name), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		structName := act.OverrideStructName
		if structName == "" {
			structName = strcase.ToCamel(act.Name)
		}

		err = actionTemplate.Execute(f, struct {
			Timestamp  time.Time
			Action     *generatedAction
			StructName string
		}{
			Timestamp:  time.Now(),
			Action:     act,
			StructName: structName,
		})
		if err != nil {
			log.Fatal(err)
		}
	}
}

var actionTemplate = template.Must(template.New("").Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated at {{ .Timestamp }}
package {{ .Action.Package }}

import (
	"fmt"
	{{- if not .Action.SkipRunPreActions }}
	"github.com/kamushadenes/chloe/errors"
	{{- else }}
	{{- if not .Action.SkipRunPostActions }}
	"github.com/kamushadenes/chloe/errors"
	{{- end }}
	{{- end }}
	{{ if not .Action.SkipSetMessage -}}"github.com/kamushadenes/chloe/langchain/memory"{{- end }}
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/langchain/actions/functions"
)

type {{ .StructName }}Action struct {
	Name        string
	Description string
	Params *action_structs.ActionParameterSet
	Extra map[string]interface{}
}

{{- if not .Action.SkipNewAction }}
// New{{ .StructName }}Action creates a new {{ .StructName }}Action with Params initialized
func New{{ .StructName }}Action() action_structs.Action {
	var params action_structs.ActionParameterSet
	
	{{- range .Action.Params }}
	params.AddParam(&action_structs.ActionParameter{
		Name:        "{{ .Name }}",
		Description: "{{ .Description }}",
		Type:        "{{ .Type }}",
		Required:    {{ .Required }},
		Enum:        {{ if .Enum }}{{ .Enum }}{{ else }}nil{{ end }},
	})

	{{- end }}

	return &{{ .StructName }}Action{
		Name:   "{{ .Action.Name }}",
		Description: "{{ .Action.Description }}",
		Params: &params,
	}
}
{{- end }}

// CheckRequiredParams checks if all required params are set
func (a *{{ .StructName }}Action) CheckRequiredParams() error {
	return a.Params.CheckRequiredParams()
}

{{- if not .Action.SkipGetName }}
func (a *{{ .StructName }}Action) GetName() string {
	return a.Name
}
{{- end }}

{{- if not .Action.SkipGetDescription }}
func (a *{{ .StructName }}Action) GetDescription() string {
	return a.Description
}
{{- end }}

{{- if not .Action.SkipSetParam }}
func (a *{{ .StructName }}Action) SetParam(key, value string) {
	a.Params.SetParam(key, value)
}
{{- end }}

{{- if not .Action.SkipGetParam }}
func (a *{{ .StructName }}Action) GetParam(key string) (string, error) {
	p, err := a.Params.GetParam(key)
	
	if p == nil {
		return "", fmt.Errorf("param %s not found", key)
	}	

	return p.Value, err
}

func (a *{{ .StructName }}Action) MustGetParam(key string) string {
	v, _ := a.GetParam(key)
	return v
}
{{- end }}

{{- if not .Action.SkipGetParams }}
func (a *{{ .StructName }}Action) GetParams() []*action_structs.ActionParameter {
	return a.Params.GetParams()
}
{{- end }}

{{- if not .Action.SkipSetMessage }}
func (a *{{ .StructName }}Action) SetMessage(msg *memory.Message) {}
{{- end }}

{{- if not .Action.SkipRunPreActions }}
func (a *{{ .StructName }}Action) RunPreActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
{{- end }}

{{- if not .Action.SkipRunPostActions }}
func (a *{{ .StructName }}Action) RunPostActions(request *action_structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
{{- end }}

func (a *{{ .StructName }}Action) GetSchema() *functions.FunctionDefinition {
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
`))
