package common

import "context"

type ASR interface {
	Transcribe(...ASRMessage) (ASRResult, error)
	TranscribeWithContext(context.Context, ...ASRMessage) (ASRResult, error)
	TranscribeWithOptions(context.Context, ASROptions) (ASRResult, error)

	Translate(...ASRMessage) (ASRResult, error)
	TranslateWithContext(context.Context, ...ASRMessage) (ASRResult, error)
	TranslateWithOptions(context.Context, ASROptions) (ASRResult, error)
}
