package structs

import (
	"context"

	"github.com/kamushadenes/chloe/langchain/memory"
	"github.com/kamushadenes/chloe/logging"
	"github.com/kamushadenes/chloe/structs/writer_structs"
	"github.com/rs/zerolog"
)

type Request interface {
	GetID() string
	GetContext() context.Context

	GetWriter() writer_structs.ChloeWriter
	SetWriter(writer_structs.ChloeWriter)

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

	GetWriter() writer_structs.ChloeWriter
	SetWriter(writer_structs.ChloeWriter)

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
