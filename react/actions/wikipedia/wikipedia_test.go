package wikipedia

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWikipediaAction(t *testing.T) {
	req := structs.NewActionRequest()
	req.Context = context.Background()
	req.Action = "wikipedia"
	req.Params["query"] = "Barack Obama"

	act := NewWikipediaAction()
	for k := range req.Params {
		act.SetParam(k, req.Params[k])
	}
	_, err := act.Execute(req)

	assert.ErrorIs(t, err, errors.ErrProceed)
}
