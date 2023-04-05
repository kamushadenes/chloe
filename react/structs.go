package react

import "github.com/rs/zerolog"

type DetectedAction struct {
	Command struct {
		Name   string `json:"name"`
		Params string `json:"params"`
	} `json:"command"`
	Thoughts struct {
		ChainOfThought []string `json:"chain_of_thought"`
		Plan           []string `json:"plan"`
		Criticism      string   `json:"criticism"`
	} `json:"thoughts"`
}

func (d DetectedAction) MarshalZerologObject(e *zerolog.Event) {
	e.Str("command_name", d.Command.Name)
	e.Str("command_params", d.Command.Params)
	e.Strs("thoughts_plan", d.Thoughts.Plan)
	e.Str("thoughts_criticism", d.Thoughts.Criticism)
	e.Strs("thoughts_chain_of_thought", d.Thoughts.ChainOfThought)
}
