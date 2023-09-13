package google

import (
	"context"
	"testing"

	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/stretchr/testify/assert"
)

func TestGoogleAction(t *testing.T) {
	req := action_structs.NewActionRequest()
	req.Context = context.Background()
	req.Action = "openai"
	req.Params["query"] = "Barack Obama"

	act := NewGoogleAction()
	for k := range req.Params {
		act.SetParam(k, req.Params[k])
	}
	_, err := act.Execute(req)

	assert.ErrorIs(t, err, errors.ErrProceed)
}
