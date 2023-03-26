package react

import (
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
	"io"
	"strings"
)

func StartAndWait(req structs.Request) {
	if req.GetStartChannel() != nil {
		req.GetStartChannel() <- true
	}
	if req.GetContinueChannel() != nil {
		<-req.GetContinueChannel()
	}
}

func NotifyError(req structs.Request, err error) error {
	logger := logging.GetLogger().With().Str("requestID", req.GetID()).Logger()
	if err != nil {
		logger.Error().Err(err).Msg("an error occurred")
	}
	if req.GetErrorChannel() != nil {
		req.GetErrorChannel() <- err
	}

	return err
}

func NotifyAndClose(req structs.Request, writer io.WriteCloser, err error) error {
	if !req.GetSkipClose() {
		if err2 := writer.Close(); err2 != nil {
			return NotifyError(req, err2)
		}
	}
	return NotifyError(req, err)
}

func WriteResult(req structs.Request, result interface{}) {
	if req.GetResultChannel() != nil {
		req.GetResultChannel() <- result
	}
}

func Truncate(s string, n int) string {
	s = strings.Join(strings.Fields(s), " ")

	if len(s) > n {
		return s[:n]
	}
	return s
}
