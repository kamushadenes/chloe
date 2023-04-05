package errors

import (
	"fmt"
)

var ErrActionFailed = fmt.Errorf("action failed")
var ErrCompletionFailed = fmt.Errorf("completion failed")
var ErrTTSFailed = fmt.Errorf("tts failed")
var ErrGenerationFailed = fmt.Errorf("generation failed")
var ErrTranscriptionFailed = fmt.Errorf("transcription failed")
var ErrModerationFailed = fmt.Errorf("moderation failed")
var ErrSummarizationFailed = Wrap(ErrCompletionFailed, fmt.Errorf("summarization failed"))
var ErrInvalidAction = fmt.Errorf("invalid action")
