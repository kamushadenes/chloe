package google

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGoogleAction(t *testing.T) {
	req := structs.NewActionRequest()
	req.Context = context.Background()
	req.Action = "google"
	req.Params = "Barack Obama"

	b := utils.BytesWriter{}
	req.Writers = append(req.Writers, &b)

	act := NewGoogleAction()
	act.SetParams(req.Params)
	act.SetWriters(req.Writers)
	err := act.Execute(req)

	assert.ErrorIs(t, err, errors.ErrProceed)
}
