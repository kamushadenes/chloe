package structs

import (
	"github.com/kamushadenes/chloe/memory"
)

type Action interface {
	GetName() string

	GetNotification() string

	SetParams(string)
	GetParams() string

	Execute(*ActionRequest) ([]*ResponseObject, error)

	SetMessage(*memory.Message)

	RunPreActions(*ActionRequest) error
	RunPostActions(*ActionRequest) error
}
