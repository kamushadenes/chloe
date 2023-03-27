package actions

import (
	structs2 "github.com/kamushadenes/chloe/react/actions/structs"
	"github.com/kamushadenes/chloe/react/utils"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

func defaultPreActions(a structs2.Action, request *structs.ActionRequest) error {
	b := &utils.BytesWriter{}

	var ws []io.WriteCloser
	ws = append(ws, request.GetWriters()...)
	ws = append(ws, b)

	a.SetWriters(ws)

	return nil
}

func defaultPostActions(a structs2.Action, request *structs.ActionRequest) error {
	truncateTokenCount := utils.GetTokenCount(request)

	for _, w := range a.GetWriters() {
		switch b := w.(type) {
		case *utils.BytesWriter:
			if err := utils.StoreChainOfThoughtResult(request, utils.Truncate(string(b.Bytes), truncateTokenCount)); err != nil {
				return err
			}
		}
	}

	return nil
}