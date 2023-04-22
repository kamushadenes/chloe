package prompts

import (
	"bytes"
	"embed"
	"fmt"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
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
		return "", errors.Wrap(errors.ErrPromptError, err)
	}

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, args); err != nil {
		return "", errors.Wrap(errors.ErrPromptError, err)
	}

	return buf.String(), nil
}

func GetExamples(prompt string, args *PromptArgs) ([]openai.ChatCompletionMessage, error) {
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
	prompt, err := GetPrompt(prompt, &PromptArgs{Args: make(map[string]interface{}), Mode: prompt})
	if err != nil {
		return 0, err
	}

	model := config.OpenAI.GetTokenizerModel(config.OpenAI.GetModel(config.Completion))

	return model.CountTokens(prompt), nil
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
