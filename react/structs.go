package react

import (
	"encoding/json"
	"github.com/rs/zerolog"
)

type DetectedAction struct {
	Command struct {
		Name   string            `json:"name"`
		Params map[string]string `json:"params"`
	} `json:"command"`
	Thoughts struct {
		ChainOfThought []string `json:"chain_of_thought"`
		Plan           []string `json:"plan"`
		Criticism      string   `json:"criticism"`
	} `json:"thoughts"`
}

func (d DetectedAction) MarshalZerologObject(e *zerolog.Event) {
	paramsJson, _ := json.Marshal(d.Command.Params)

	e.Str("command_name", d.Command.Name)
	e.RawJSON("command_params", paramsJson)
	e.Strs("thoughts_plan", d.Thoughts.Plan)
	e.Str("thoughts_criticism", d.Thoughts.Criticism)
	e.Strs("thoughts_chain_of_thought", d.Thoughts.ChainOfThought)
}
