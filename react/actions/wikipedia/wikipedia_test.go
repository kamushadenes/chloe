package wikipedia

import (
	"context"
	"github.com/kamushadenes/chloe/react/errors"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWikipediaAction(t *testing.T) {
	req := structs.NewActionRequest()
	req.Context = context.Background()
	req.Action = "wikipedia"
	req.Params = "Barack Obama"

	b := utils.BytesWriter{}
	req.Writers = append(req.Writers, &b)

	act := NewWikipediaAction()
	act.SetParams(req.Params)
	act.SetWriters(req.Writers)
	err := act.Execute(req)

	assert.ErrorIs(t, err, errors.ErrProceed)
}
