package structs

import (
	"context"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
	"io"
)

type Request interface {
	GetID() string
	GetContext() context.Context
	GetWriters() []io.WriteCloser
	GetSkipClose() bool
	GetStartChannel() chan bool
	GetContinueChannel() chan bool
	GetErrorChannel() chan error
	GetResultChannel() chan interface{}
	GetMessage() *memory.Message
}

type ImageRequest interface {
	Request
	GetSize() string
}

func LoggerFromRequest(request Request) zerolog.Logger {
	return logging.GetLogger().With().Str("requestID", request.GetID()).Logger()
}
