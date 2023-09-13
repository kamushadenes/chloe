package functions

import (
	"github.com/kamushadenes/chloe/structs/response_object_structs"
)

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

func (fnc *FunctionCall) Run() ([]*response_object_structs.ResponseObject, error) {
	/*
		w := writer_structs.NewMockWriter()

		req := action_structs.NewActionRequest()

		req.Action = fnc.Name
		req.Writer = w

		var args map[string]string
		if err := json.Unmarshal([]byte(fnc.Arguments), &args); err != nil {
			return nil, err
		}

		req.Params = args

		if err := actions.HandleAction(req); err != nil {
			return nil, err
		}

		return w.GetObjects(), nil
	*/

	return nil, nil
}
