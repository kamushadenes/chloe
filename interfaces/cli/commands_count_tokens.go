package cli

import (
	"fmt"
	"github.com/kamushadenes/chloe/tokenizer"
	"strings"
)

type CountTokensCmd struct {
	Prompt []string `arg:"" help:"Prompt to generate"`
	Model  string   `help:"Model to use" default:"gpt-3.5-turbo"`
}

func (c *CountTokensCmd) Run(globals *Globals) error {
	fmt.Println(tokenizer.CountTokens(c.Model, strings.Join(c.Prompt, " ")))

	return nil
}
