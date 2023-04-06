package actions

import (
	"github.com/kamushadenes/chloe/structs"
)

func defaultPreActions(a structs.Action, request *structs.ActionRequest) error {
	return nil
}

func defaultPostActions(a structs.Action, request *structs.ActionRequest) error {
	return nil
}
