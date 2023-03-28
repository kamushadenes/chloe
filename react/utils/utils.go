package utils

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"github.com/kamushadenes/chloe/config"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
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
	go func() {
		if req.GetErrorChannel() != nil {
			req.GetErrorChannel() <- err
		}
	}()

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

func StoreChainOfThoughtResult(request structs.ActionOrCompletionRequest, content string) error {
	nmsg := memory.NewMessage(uuid.Must(uuid.NewV4()).String(), request.GetMessage().Interface)
	nmsg.Role = "user"

	params := struct {
		Result string `json:"result"`
	}{
		Result: content,
	}

	b, err := json.Marshal(params)
	if err != nil {
		return err
	}

	nmsg.SetContent(string(b))
	nmsg.User = request.GetMessage().User

	return nmsg.Save(request.GetContext())
}

func GetAvailableTokenCount(request *structs.ActionRequest) int {
	tokenCount := request.CountTokens()

	maxTokens := config.OpenAI.GetMaxTokens(config.OpenAI.GetModel(config.ChainOfThought))

	truncateTokenCount := maxTokens - tokenCount - config.OpenAI.GetMinReplyTokens()

	if truncateTokenCount < config.OpenAI.GetMinReplyTokens() {
		truncateTokenCount = config.OpenAI.GetMinReplyTokens()
	}

	return truncateTokenCount
}
