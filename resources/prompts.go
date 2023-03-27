package resources

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"strings"
	"text/template"
)

//go:embed prompts/chatgpt/*.prompt
var prompts embed.FS

type PromptArgs struct {
	Args map[string]interface{} `json:"args"`
	Mode string                 `json:"mode"`
}

func GetPrompt(prompt string, args *PromptArgs) (string, error) {
	tmpl, err := template.ParseFS(
		prompts,
		fmt.Sprintf("prompts/chatgpt/%s.prompt", prompt),
	)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, args); err != nil {
		return "", err
	}

	return buf.String(), nil
}

func GetExamples(prompt string, args *PromptArgs) ([]openai.ChatCompletionMessage, error) {
	tmpl, err := template.ParseFS(
		prompts,
		fmt.Sprintf("prompts/chatgpt/%s.prompt.examples", prompt),
	)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, args); err != nil {
		return nil, err
	}

	var examples []openai.ChatCompletionMessage

	for _, line := range strings.Split(buf.String(), "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 2 {
			name := fields[0]
			example := strings.Join(fields[1:], " ")

			examples = append(examples, openai.ChatCompletionMessage{
				Content: example,
				Role:    "system",
				Name:    name,
			})
		}
	}

	return examples, nil
}

func GetPromptSize(prompt string) (int, error) {
	tmpl, err := template.ParseFS(
		prompts,
		fmt.Sprintf("prompts/chatgpt/%s.prompt", prompt),
	)
	if err != nil {
		return 0, err
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, nil); err != nil {
		return 0, err
	}

	return len(strings.Fields(buf.String())), nil
}

func ListPrompts() ([]string, error) {
	// lists all files that end with .prompt in prompts/chatgpt
	// and returns the file names without the .prompt extension
	entries, err := prompts.ReadDir("prompts/chatgpt")
	if err != nil {
		return nil, err
	}

	var prompts []string
	for _, entry := range entries {
		prompts = append(prompts, strings.TrimSuffix(entry.Name(), ".prompt"))
	}

	return prompts, nil
}
