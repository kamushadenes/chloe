package react

import (
	"context"
	"github.com/kamushadenes/chloe/memory"
	"github.com/kamushadenes/chloe/structs"
	"io"
)

type Action interface {
	GetName() string
	GetNotification() string
	SetParams(string)
	GetParams() string
	Execute(context.Context) error
	SetWriters([]io.WriteCloser)
	GetWriters() []io.WriteCloser
	SetMessage(*memory.Message)
	RunPreActions(*structs.ActionRequest) error
	RunPostActions(*structs.ActionRequest) error
}
