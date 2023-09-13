package cli

import (
	"github.com/kamushadenes/chloe/flags"
	"github.com/kamushadenes/chloe/langchain/tts/common"
	"github.com/kamushadenes/chloe/langchain/tts/google"
	"strings"
)

type TTSCmd struct {
	Prompt     []string `arg:"" help:"Prompt to generate"`
	OutputPath string   `help:"Output path, if not specified, it will be printed to stdout if not a TTY, otherwise it will be saved to generated.mp3" type:"path"`
}

func (c *TTSCmd) Run(globals *Globals) error {
	tts := google.NewTTSGoogle()

	res, err := tts.TTSWithContext(globals.Context, common.TTSMessage{Text: strings.Join(c.Prompt, " ")})
	if err != nil {
		return err
	}

	if len(c.OutputPath) > 0 {
		_, err = NewFileWriter(c.OutputPath).Write(res.Audio)
		return err
	}

	if flags.InteractiveCLI {
		_, err = NewFileWriter("generated.mp3").Write(res.Audio)
		return err
	}

	_, err = NewCLIWriter().Write(res.Audio)
	return err
}
