package functions

import "encoding/json"

type FunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments,omitempty"`
}

type FunctionDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	// Parameters is an object describing the function.
	// You can pass json.RawMessage to describe the schema,
	// or you can pass in a struct which serializes to the proper JSON schema.
	// The jsonschema package is provided for convenience, but you should
	// consider another specialized library if you require more complex schemas.
	Parameters any `json:"parameters"`
}

type FunctionResult struct {
	Result string `json:"result"`
}

func (f *FunctionResult) String() string {
	b, _ := json.Marshal(f)

	return string(b)
}
