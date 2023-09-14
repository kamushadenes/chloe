package common

import "context"

type ASR interface {
	Transcribe(string) (ASRResult, error)
	TranscribeWithContext(context.Context, string) (ASRResult, error)
	TranscribeWithOptions(context.Context, ASROptions) (ASRResult, error)

	Translate(string) (ASRResult, error)
	TranslateWithContext(context.Context, string) (ASRResult, error)
	TranslateWithOptions(context.Context, ASROptions) (ASRResult, error)
}
