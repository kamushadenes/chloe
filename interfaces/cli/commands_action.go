package cli

import (
	"fmt"
	"github.com/kamushadenes/chloe/channels"
	"github.com/kamushadenes/chloe/structs"
	"strings"
)

type ActionCmd struct {
	Action string   `arg:"" help:"Action to perform" enum:"google,image,latex,math,news,scrape,transcribe,tts,wikipedia,youtube_summarizer"`
	Params []string `arg:"" help:"Parameters for the action"`
}

func (a *ActionCmd) Run(globals *Globals) error {
	req := structs.NewActionRequest()
	req.Context = globals.Context
	req.Action = a.Action
	req.Params = strings.Join(a.Params, " ")
	req.Thought = fmt.Sprintf("User wants to run action %s", a.Action)
	req.Writer = NewCLIWriter()

	return channels.RunAction(req)
}
