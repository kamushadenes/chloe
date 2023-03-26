package react

import (
	"github.com/kamushadenes/chloe/structs"
	"io"
)

func defaultPreActions(a Action, request *structs.ActionRequest) error {
	b := &BytesWriter{}

	var ws []io.WriteCloser
	ws = append(ws, request.GetWriters()...)
	ws = append(ws, b)

	a.SetWriters(ws)

	return nil
}

func defaultPostActions(a Action, request *structs.ActionRequest) error {
	truncateTokenCount := getTokenCount(request)

	for _, w := range a.GetWriters() {
		switch b := w.(type) {
		case *BytesWriter:
			if err := storeChainOfThoughtResult(request, Truncate(string(b.Bytes), truncateTokenCount)); err != nil {
				return err
			}
		}
	}

	return nil
}
