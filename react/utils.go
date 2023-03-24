package react

import (
	"github.com/kamushadenes/chloe/structs"
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
	if req.GetErrorChannel() != nil {
		req.GetErrorChannel() <- err
	}

	return err
}

func NotifyAndClose(req structs.Request, err error) error {
	for _, writer := range req.GetWriters() {
		if !req.GetSkipClose() {
			err := writer.Close()
			_ = NotifyError(req, err)
		} else {
			_ = NotifyError(req, err)
		}
	}

	return err
}

func WriteResult(req structs.Request, result interface{}) {
	if req.GetResultChannel() != nil {
		req.GetResultChannel() <- result
	}
}

func Truncate(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}
