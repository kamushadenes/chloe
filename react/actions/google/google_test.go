package google

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleAction(t *testing.T) {
	req := structs.NewActionRequest()
	req.Context = context.Background()
	req.Action = "google"
	req.Params = "Barack Obama"

	act := NewGoogleAction()
	act.SetParams(req.Params)
	_, err := act.Execute(req)

	assert.ErrorIs(t, err, errors.ErrProceed)
}
