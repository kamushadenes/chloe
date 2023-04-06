package channels

import (
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
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
