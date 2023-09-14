package actions

import (
	"context"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/structs/action_structs"
	"github.com/kamushadenes/chloe/structs/writer_structs"
	"github.com/stretchr/testify/assert"
)

func TestActions_Metadata(t *testing.T) {
	for k := range actions {
		k2 := k
		t.Run(k, func(t *testing.T) {
			t.Parallel()

			act := actions[k2]()

			assert.True(t, len(act.GetName()) > 0)
			assert.True(t, len(act.GetNotification()) > 0)

			assert.True(t, len(act.GetParams()) > 0)
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
		req := action_structs.NewActionRequest()
		req.Context = context.Background()
		req.Params["foo"] = tt.params
		req.Action = "mock"
		req.Writer = writer_structs.NewMockWriter()
		req.Message = memory.NewMessage(uuid.Must(uuid.NewV4()).String(), "test")

		err := HandleAction(req)
		if tt.wantErr {
			assert.ErrorIs(t, err, errors.ErrMock)
		} else {
			assert.NoError(t, err)
		}
	}
}
