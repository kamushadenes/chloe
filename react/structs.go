package react

import "github.com/rs/zerolog"

type DetectedAction struct {
	Command struct {
		Name   string `json:"name"`
		Params string `json:"params"`
	} `json:"command"`
	Thoughts struct {
		Text      string   `json:"text"`
		Reasoning string   `json:"reasoning"`
		Plan      []string `json:"plan"`
		Criticism string   `json:"criticism"`
		Speak     string   `json:"speak"`
	} `json:"thoughts"`
}

func (d DetectedAction) MarshalZerologObject(e *zerolog.Event) {
	e.Str("command_name", d.Command.Name)
	e.Str("command_params", d.Command.Params)
	e.Str("thoughts_text", d.Thoughts.Text)
	e.Str("thoughts_reasoning", d.Thoughts.Reasoning)
	e.Strs("thoughts_plan", d.Thoughts.Plan)
	e.Str("thoughts_criticism", d.Thoughts.Criticism)
	e.Str("thoughts_speak", d.Thoughts.Speak)
}
