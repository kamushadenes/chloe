package common

import "context"

type TTS interface {
	TTS(TTSMessage) (TTSResult, error)
	TTSWithContext(context.Context, TTSMessage) (TTSResult, error)
	TTSWithOptions(context.Context, TTSOptions) (TTSResult, error)
}
