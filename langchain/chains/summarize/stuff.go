package summarize

import (
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/langchain/llm/base"
	"github.com/kamushadenes/chloe/langchain/prompts"
)

func Stuff(text string) (string, error) {
	prompt, err := prompts.GetPrompt("summarize_stuff", struct{ Text string }{text})
	if err != nil {
		return "", err
	}

	llm := base.NewLLMWithDefaultModel(config.LLM.Provider)

	res, err := llm.Generate(prompt)
	if err != nil {
		return "", err
	}

	return res.Generations[0].Text, nil
}
