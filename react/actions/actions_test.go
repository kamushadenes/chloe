package actions

import (
	"context"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestActions_Metadata(t *testing.T) {
	for k := range actions {
		t.Run(k, func(t *testing.T) {
			t.Parallel()

			act := actions[k]()

			assert.True(t, len(act.GetName()) > 0)
			assert.True(t, len(act.GetNotification()) > 0)

			assert.True(t, len(act.GetParams()) == 0)
			act.SetParam("test", "foo")
			p, err := act.GetParam("test")
			assert.NoError(t, err)
			assert.True(t, p == "foo")
		})
	}
}

func TestActions_HandleAction(t *testing.T) {
	tests := []struct {
		name    string
		params  string
		wantErr bool
	}{
		{
			name:    "mock",
			params:  "test",
			wantErr: false,
		},
		{
			name:    "mockErrExecute",
			params:  "err",
			wantErr: true,
		},
		{
			name:    "mockErrPreActions",
			params:  "errPre",
			wantErr: true,
		},
		{
			name:    "mockErrPostActions",
			params:  "errPost",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		req := structs.NewActionRequest()
		req.Context = context.Background()
		req.Params["foo"] = tt.params
		req.Action = "mock"
		req.Writer = structs.NewMockWriter()
		req.Message = memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "test")

		err := HandleAction(req)
		if tt.wantErr {
			assert.ErrorIs(t, err, errors.ErrMock)
		} else {
			assert.NoError(t, err)
		}
	}
}
