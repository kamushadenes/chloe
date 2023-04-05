package cli

import (
	"github.com/kamushadenes/chloe/flags"
	"os"
	"strings"
)

type GenerateCmd struct {
	Prompt     []string `arg:"" help:"Prompt to generate"`
	OutputPath string   `help:"Output path, if not specified, it will be printed to stdout if not a TTY, otherwise it will be saved to generated.png" type:"path"`
}

func (c *GenerateCmd) Run(globals *Globals) error {
	if len(c.OutputPath) > 0 {
		f, err := os.Create(c.OutputPath)
		if err != nil {
			return err
		}

		return Generate(globals.Context, strings.Join(c.Prompt, " "), f)
	}

	if flags.InteractiveCLI {
		f, err := os.Create("generated.png")
		if err != nil {
			return err
		}
		return Generate(globals.Context, strings.Join(c.Prompt, " "), f)
	}

	return Generate(globals.Context, strings.Join(c.Prompt, " "))
}
