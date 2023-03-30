package cli

import (
	"github.com/kamushadenes/chloe/flags"
	"os"
	"strings"
)

type TTSCmd struct {
	Prompt     []string `arg:"" help:"Prompt to generate"`
	OutputPath string   `help:"Output path, if not specified, it will be printed to stdout if not a TTY, otherwise it will be saved to generated.mp3" type:"path"`
}

func (c *TTSCmd) Run(globals *Globals) error {
	if len(c.OutputPath) > 0 {
		f, err := os.Create(c.OutputPath)
		if err != nil {
			return err
		}

		return TTS(globals.Context, strings.Join(c.Prompt, " "), f)
	}

	if flags.InteractiveCLI {
		f, err := os.Create("generated.mp3")
		if err != nil {
			return err
		}
		return TTS(globals.Context, strings.Join(c.Prompt, " "), f)
	}
	return TTS(globals.Context, strings.Join(c.Prompt, " "))
}
