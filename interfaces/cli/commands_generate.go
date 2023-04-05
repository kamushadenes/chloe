package cli

import (
	"github.com/kamushadenes/chloe/flags"
	"strings"
)

type GenerateCmd struct {
	Prompt     []string `arg:"" help:"Prompt to generate"`
	OutputPath string   `help:"Output path, if not specified, it will be printed to stdout if not a TTY, otherwise it will be saved to generated.png" type:"path"`
}

func (c *GenerateCmd) Run(globals *Globals) error {
	if len(c.OutputPath) > 0 {
		return Generate(globals.Context, strings.Join(c.Prompt, " "), NewFileWriter(c.OutputPath))
	}

	if flags.InteractiveCLI {
		return Generate(globals.Context, strings.Join(c.Prompt, " "), NewFileWriter("generated.png"))
	}

	return Generate(globals.Context, strings.Join(c.Prompt, " "), NewCLIWriter())
}
