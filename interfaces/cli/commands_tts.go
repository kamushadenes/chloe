package cli

import (
	"github.com/kamushadenes/chloe/flags"
	"strings"
)

type TTSCmd struct {
	Prompt     []string `arg:"" help:"Prompt to generate"`
	OutputPath string   `help:"Output path, if not specified, it will be printed to stdout if not a TTY, otherwise it will be saved to generated.mp3" type:"path"`
}

func (c *TTSCmd) Run(globals *Globals) error {
	if len(c.OutputPath) > 0 {
		return TTS(globals.Context, strings.Join(c.Prompt, " "), NewFileWriter(c.OutputPath))
	}

	if flags.InteractiveCLI {
		return TTS(globals.Context, strings.Join(c.Prompt, " "), NewFileWriter("generated.mp3"))
	}

	return TTS(globals.Context, strings.Join(c.Prompt, " "), NewCLIWriter())
}
