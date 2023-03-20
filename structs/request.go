package structs

import (
	"context"
	"io"
)

type Request interface {
	GetContext() context.Context
	GetWriters() []io.WriteCloser
	GetSkipClose() bool
	GetStartChannel() chan bool
	GetContinueChannel() chan bool
	GetErrorChannel() chan error
	GetResultChannel() chan interface{}
}
