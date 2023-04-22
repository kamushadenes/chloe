package actions

import (
	"bytes"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
)

func StoreActionDetectionResult(request structs.ActionOrCompletionRequest, role string, content string, summary string) error {
	logger := logging.FromContext(request.GetContext())

	logger.Debug().Str("content", content).Msg("storing action detection result")

	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.GetMessage().Interface)
	nmsg.Role = role
	nmsg.Summary = summary

	params := struct {
		Result string `json:"result"`
	}{
		Result: content,
	}

	b, err := json.Marshal(params)
	if err != nil {
		return errors.Wrap(errors.ErrSaveMessage, err)
	}
	var buf bytes.Buffer
	err = json.Compact(&buf, b)
	if err != nil {
		return errors.Wrap(errors.ErrSaveMessage, err)
	}

	nmsg.SetContent(buf.String())
	nmsg.User = request.GetMessage().User

	return nmsg.Save(request.GetContext())
}
