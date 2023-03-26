package react

import (
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

func scrapePreActions(a Action, request *structs.ActionRequest) error {
	logger := logging.GetLogger().With().Str("action", a.GetName()).Str("params", a.GetParams()).Logger()

	truncateTokenCount := getTokenCount(request)

	b := &BytesWriter{}
	a.SetWriters([]io.WriteCloser{b})

	logger.Info().Msg("executing action")
	err := a.Execute(request.Context)
	if err != nil {
		return err
	}

	if err := storeChainOfThoughtResult(request, Truncate(string(b.Bytes), truncateTokenCount)); err != nil {
		return err
	}
	return ErrProceed
}
