package actions

import "github.com/kamushadenes/chloe/structs/action_structs"

func defaultPreActions(a action_structs.Action, request *action_structs.ActionRequest) error {
	return nil
}

func defaultPostActions(a action_structs.Action, request *action_structs.ActionRequest) error {
	return nil
}
