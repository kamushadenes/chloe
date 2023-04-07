//go:build ignore

package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"log"
	"os"
	"text/template"
	"time"
)

type generatedAction struct {
	Name               string
	OverrideStructName string
	Package            string
	RequiredParams     []string
	SkipNewAction      bool
	SkipGetName        bool
	SkipSetParam       bool
	SkipGetParam       bool
	SkipGetParams      bool
	SkipRunPreActions  bool
	SkipRunPostActions bool
	SkipSetMessage     bool
}

var actionsToGenerate = []*generatedAction{
	{
		Name:           "append_file",
		Package:        "file",
		RequiredParams: []string{"path", "content"},
	},
	{
		Name:           "delete_file",
		Package:        "file",
		RequiredParams: []string{"path"},
	},
	{
		Name:           "read_file",
		Package:        "file",
		RequiredParams: []string{"path"},
	},
	{
		Name:           "write_file",
		Package:        "file",
		RequiredParams: []string{"path", "content"},
	},
	{
		Name:               "google",
		Package:            "google",
		RequiredParams:     []string{"query"},
		SkipRunPostActions: true,
	},
	{
		Name:              "image",
		Package:           "image",
		RequiredParams:    []string{"prompt"},
		SkipRunPreActions: true,
	},
	{
		Name:           "variation",
		Package:        "image",
		RequiredParams: []string{"path"},
	},
	{
		Name:           "latex",
		Package:        "latex",
		RequiredParams: []string{"formula"},
	},
	{
		Name:           "math",
		Package:        "math",
		RequiredParams: []string{"expression"},
	},
	{
		Name:           "midjourney_prompt_generator",
		Package:        "midjourney_prompt_generator",
		RequiredParams: []string{"prompt"},
	},
	{
		Name:               "mock",
		Package:            "mock",
		RequiredParams:     []string{"foo"},
		SkipRunPreActions:  true,
		SkipRunPostActions: true,
	},
	{
		Name:           "news",
		Package:        "news",
		RequiredParams: []string{"query"},
	},
	{
		Name:           "news_by_country",
		Package:        "news",
		RequiredParams: []string{"query"},
	},
	{
		Name:               "scrape",
		Package:            "scrape",
		RequiredParams:     []string{"url"},
		SkipRunPreActions:  true,
		SkipRunPostActions: true,
	},
	{
		Name:               "transcribe",
		Package:            "transcribe",
		RequiredParams:     []string{"path"},
		SkipRunPostActions: true,
		SkipSetMessage:     true,
	},
	{
		Name:               "tts",
		Package:            "tts",
		OverrideStructName: "TTS",
		RequiredParams:     []string{"text"},
	},
	{
		Name:               "wikipedia",
		Package:            "wikipedia",
		RequiredParams:     []string{"query"},
		SkipRunPostActions: true,
	},
	{
		Name:           "youtube_summarize",
		Package:        "youtube",
		RequiredParams: []string{"url"},
	},
	{
		Name:           "youtube_transcribe",
		Package:        "youtube",
		RequiredParams: []string{"url"},
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
	"github.com/kamushadenes/chloe/errors"
	{{ if not .Action.SkipSetMessage -}}"github.com/kamushadenes/chloe/memory"{{- end }}
	"github.com/kamushadenes/chloe/structs"
)


{{- if not .Action.SkipNewAction }}
// New{{ .StructName }}Action creates a new {{ .StructName }}Action with Params initialized
func New{{ .StructName }}Action() structs.Action {
	return &{{ .StructName }}Action{
		Name:   "{{ .Action.Name }}",
		Params: make(map[string]string),
	}
}
{{- end }}

// CheckRequiredParams checks if all required params are set
func (a *{{ .StructName }}Action) CheckRequiredParams() error {
	required := []string{
		{{- range .Action.RequiredParams }}
		"{{ . }}",
		{{- end }}
	}

	for k := range required {
		if _, err := a.GetParam(required[k]); err != nil {
			return errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("required param %s is not set", required[k]))
		}
	}

	return nil
}


{{- if not .Action.SkipGetName }}
func (a *{{ .StructName }}Action) GetName() string {
	return a.Name
}
{{- end }}

{{- if not .Action.SkipSetParam }}
func (a *{{ .StructName }}Action) SetParam(key, value string) {
	a.Params[key] = value
}
{{- end }}

{{- if not .Action.SkipGetParam }}
func (a *{{ .StructName }}Action) GetParam(key string) (string, error) {
	if value, ok := a.Params[key]; ok {
		return value, nil
	}

	return "", errors.Wrap(errors.ErrInvalidParameter, fmt.Errorf("param %s not found", key))
}
{{- end }}

{{- if not .Action.SkipGetParams }}
func (a *{{ .StructName }}Action) GetParams() map[string]string {
	return a.Params
}
{{- end }}

{{- if not .Action.SkipSetMessage }}
func (a *{{ .StructName }}Action) SetMessage(msg *memory.Message) {}
{{- end }}

{{- if not .Action.SkipRunPreActions }}
func (a *{{ .StructName }}Action) RunPreActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
{{- end }}

{{- if not .Action.SkipRunPostActions }}
func (a *{{ .StructName }}Action) RunPostActions(request *structs.ActionRequest) error {
	return errors.ErrNotImplemented
}
{{- end }}
`))
