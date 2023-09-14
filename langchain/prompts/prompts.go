package prompts

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/kamushadenes/chloe/tokenizer"
	"strings"
	"text/template"

	"github.com/kamushadenes/chloe/errors"
)

//go:embed prompts/chatgpt/*.prompt
var prompts embed.FS

type PromptArgs struct {
	Args map[string]interface{} `json:"args"`
	Mode string                 `json:"mode"`
}

type BootstrapArgs struct {
	Date          string
	Time          string
	Interface     string
	UserID        uint
	UserFirstName string
	UserLastName  string
}

func GetPrompt(prompt string, args interface{}) (string, error) {
	tmpl, err := template.ParseFS(
		prompts,
		fmt.Sprintf("prompts/chatgpt/%s.prompt", prompt),
	)

	if err != nil {
		return "", errors.Wrap(errors.ErrPromptError, err)
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, args); err != nil {
		return "", errors.Wrap(errors.ErrPromptError, err)
	}

	return buf.String(), nil
}

/*
func GetExamples(prompt string, args interface{}) ([]messages.Message, error) {
	tmpl, err := template.ParseFS(
		prompts,
		fmt.Sprintf("prompts/chatgpt/%s.prompt.examples", prompt),
	)
	if err != nil {
		return nil, errors.Wrap(errors.ErrPromptError, err)
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, args); err != nil {
		return nil, errors.Wrap(errors.ErrPromptError, err)
	}

	var examples []messages.Message

	for _, line := range strings.Split(buf.String(), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 2 {
			name := fields[0]
			example := strings.Join(fields[1:], " ")

			examples = append(examples, messages.Message{
				Content: example,
				Role:    messages.System,
				Name:    name,
			})
		}
	}

	return examples, nil
}
*/

func GetPromptSize(prompt string) (int, error) {
	prompt, err := GetPrompt(prompt, &PromptArgs{Args: make(map[string]interface{}), Mode: prompt})
	if err != nil {
		return 0, err
	}

	return tokenizer.CountTokens("gpt-3.5-turbo", prompt), nil
}

func ListPrompts() ([]string, error) {
	entries, err := prompts.ReadDir("prompts/chatgpt")
	if err != nil {
		return nil, errors.Wrap(errors.ErrPromptError, err)
	}

	var prompts []string
	for _, entry := range entries {
		prompts = append(prompts, strings.TrimSuffix(entry.Name(), ".prompt"))
	}

	return prompts, nil
}
