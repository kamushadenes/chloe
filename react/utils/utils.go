package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"github.com/kamushadenes/chloe/utils"
	"io"
)

func StartAndWait(req structs.Request) {
	if req.GetStartChannel() != nil {
		req.GetStartChannel() <- true
	}
	if req.GetContinueChannel() != nil {
		<-req.GetContinueChannel()
	}
}

func NotifyError(req structs.Request, errs ...error) error {
	logger := logging.FromContext(req.GetContext()).With().Str("requestID", req.GetID()).Logger()

	for k := range errs {
		if errs[k] != nil {
			logger.Error().Errs("errors", errs).Msg("an error occurred")
			break
		}
	}

	go func() {
		if req.GetErrorChannel() != nil {
			req.GetErrorChannel() <- errors.Wrap(errs...)
		}
	}()

	return errors.Wrap(errs...)
}

func NotifyAndClose(req structs.Request, writer io.WriteCloser, errs ...error) error {
	if !req.GetSkipClose() {
		if err2 := writer.Close(); err2 != nil {
			// TODO: we're losing the original error here
			return NotifyError(req, errors.Wrap(errors.ErrCloseWriter, err2))
		}
	}
	return NotifyError(req, errs...)
}

func WriteResult(req structs.Request, result interface{}) {
	if req.GetResultChannel() != nil {
		req.GetResultChannel() <- result
	}
}

func StoreActionDetectionResult(request structs.ActionOrCompletionRequest, content string) error {
	logger := logging.FromContext(request.GetContext())

	logger.Debug().Str("content", content).Msg("storing action detection result")

	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.GetMessage().Interface)
	nmsg.Role = "user"

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

func GetAvailableTokenCount(request *structs.ActionRequest) int {
	return utils.SubtractIntWithMinimum(
		config.OpenAI.GetMinReplyTokens(),
		config.OpenAI.GetModel(config.ChainOfThought).GetContextSize(),
		request.CountTokens(),
		config.OpenAI.GetMinReplyTokens(),
	)
}
