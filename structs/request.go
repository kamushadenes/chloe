package structs

import (
	"context"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/memory"
	"github.com/rs/zerolog"
)

type Request interface {
	GetID() string
	GetContext() context.Context

	GetWriter() ChloeWriter
	SetWriter(ChloeWriter)

	GetSkipClose() bool

	GetStartChannel() chan bool
	GetContinueChannel() chan bool
	GetErrorChannel() chan error
	GetResultChannel() chan interface{}

	GetMessage() *memory.Message
}

type ActionOrCompletionRequest interface {
	GetID() string
	GetContext() context.Context

	GetWriter() ChloeWriter
	SetWriter(ChloeWriter)

	GetSkipClose() bool
	GetMessage() *memory.Message
}

type ImageRequest interface {
	Request
	GetSize() string
}

func LoggerFromRequest(request ActionOrCompletionRequest) zerolog.Logger {
	return logging.GetLogger().With().Str("requestID", request.GetID()).Logger()
}
