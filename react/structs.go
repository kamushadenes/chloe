package react

type ChainOfThoughtResponse struct {
	Thought string `json:"thought"`
	Action  string `json:"action"`
	Params  string `json:"params"`
}
