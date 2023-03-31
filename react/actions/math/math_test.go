package math

import (
	"context"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateAction(t *testing.T) {
	req := structs.NewActionRequest()
	req.Context = context.Background()
	req.Action = "calculate"
	req.Params = "4 * 7 / 2"

	b := utils.BytesWriter{}
	req.Writers = append(req.Writers, &b)

	act := NewCalculateAction()
	act.SetParams(req.Params)
	act.SetWriters(req.Writers)
	err := act.Execute(req)

	assert.NoError(t, err)
	assert.Equal(t, "14", string(b.Bytes))
}
