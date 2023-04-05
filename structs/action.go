package structs

import (
	"github.com/kamushadenes/chloe/memory"
	"io"
)

type Action interface {
	GetName() string
	GetNotification() string
	SetParams(string)
	GetParams() string
	Execute(*ActionRequest) error
	SetWriters([]io.WriteCloser)
	GetWriters() []io.WriteCloser
	SetMessage(*memory.Message)
	RunPreActions(*ActionRequest) error
	RunPostActions(*ActionRequest) error
}
