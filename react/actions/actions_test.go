package actions

import (
	"context"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
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
			act.SetParams("test")
			assert.True(t, act.GetParams() == "test")

			assert.True(t, len(act.GetWriters()) == 0)
			act.SetWriters([]io.WriteCloser{os.Stdout})
			assert.True(t, len(act.GetWriters()) == 1)
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
		req.Params = tt.params

		b := utils.BytesWriter{}

		req.Writers = []io.WriteCloser{&b}

		req.Action = "mock"

		err := HandleAction(req)
		if tt.wantErr {
			assert.ErrorIs(t, err, errors.ErrMock)
		} else {
			assert.NoError(t, err)
		}
	}
}
