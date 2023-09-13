package cli

import (
	"strings"

	"github.com/kamushadenes/chloe/structs/action_structs"
)

type ActionCmd struct {
	Action string   `arg:"" help:"Action to perform" enum:"openai,image,latex,math,news,scrape,transcribe,tts,wikipedia,youtube_summarizer"`
	Params []string `arg:"" help:"Parameters for the action"`
}

func (a *ActionCmd) Run(globals *Globals) error {
	req := action_structs.NewActionRequest()
	req.Context = globals.Context
	req.Action = a.Action
	req.Params["text"] = strings.Join(a.Params, " ")
	req.Writer = NewCLIWriter()

	return nil

	//return structs.RunAction(req)
}
