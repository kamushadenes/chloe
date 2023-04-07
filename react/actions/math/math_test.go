package math

import (
	"context"
	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateAction(t *testing.T) {
	req := structs.NewActionRequest()
	req.Context = context.Background()
	req.Action = "calculate"
	req.Params["expression"] = "4 * 7 / 2"

	act := NewMathAction()
	for k := range req.Params {
		act.SetParam(k, req.Params[k])
	}
	objs, err := act.Execute(req)

	assert.NoError(t, err)
	assert.Equal(t, "14", objs[0].String())
}
