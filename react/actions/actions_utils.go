package actions

import (
	"github.com/kamushadenes/chloe/structs"
)

func defaultPreActions(a structs.Action, request *structs.ActionRequest) error {
	return nil
}

func defaultPostActions(a structs.Action, request *structs.ActionRequest) error {
	/*
		truncateTokenCount := utils.GetAvailableTokenCount(request)

		for _, w := range a.GetWriters() {
			switch b := w.(type) {
			case *utils.BytesWriter:
				if err := utils.StoreActionDetectionResult(request, utils2.Truncate(string(b.Bytes), truncateTokenCount)); err != nil {
					return err
				}
			}
		}
	*/

	return nil
}
