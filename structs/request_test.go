package structs

import (
	"context"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestRequest(t *testing.T) {
	ctx := context.Background()
	w := io.WriteCloser(nil)
	skipClose := true
	startChannel := make(chan bool)
	continueChannel := make(chan bool)
	errorChannel := make(chan error)
	resultChannel := make(chan interface{})

	tests := []struct {
		name string
		s    Request
	}{
		{
			name: "Calculation",
			s: &CalculationRequest{
				Context:         ctx,
				Writer:          w,
				SkipClose:       skipClose,
				StartChannel:    startChannel,
				ContinueChannel: continueChannel,
				ErrorChannel:    errorChannel,
				ResultChannel:   resultChannel,
			},
		},
		{
			name: "Completion",
			s: &CompletionRequest{
				Context:         ctx,
				Writer:          w,
				SkipClose:       skipClose,
				StartChannel:    startChannel,
				ContinueChannel: continueChannel,
				ErrorChannel:    errorChannel,
				ResultChannel:   resultChannel,
			},
		},
		{
			name: "Generation",
			s: &GenerationRequest{
				Context:         ctx,
				Writers:         []io.WriteCloser{w},
				SkipClose:       skipClose,
				StartChannel:    startChannel,
				ContinueChannel: continueChannel,
				ErrorChannel:    errorChannel,
				ResultChannel:   resultChannel,
			},
		},
		{
			name: "Scrape",
			s: &ScrapeRequest{
				Context:         ctx,
				Writer:          w,
				SkipClose:       skipClose,
				StartChannel:    startChannel,
				ContinueChannel: continueChannel,
				ErrorChannel:    errorChannel,
				ResultChannel:   resultChannel,
			},
		},
		{
			name: "Transcription",
			s: &TranscriptionRequest{
				Context:         ctx,
				Writer:          w,
				SkipClose:       skipClose,
				StartChannel:    startChannel,
				ContinueChannel: continueChannel,
				ErrorChannel:    errorChannel,
				ResultChannel:   resultChannel,
			},
		},
		{
			name: "TTS",
			s: &TTSRequest{
				Context:         ctx,
				Writers:         []io.WriteCloser{w},
				SkipClose:       skipClose,
				StartChannel:    startChannel,
				ContinueChannel: continueChannel,
				ErrorChannel:    errorChannel,
				ResultChannel:   resultChannel,
			},
		},
		{
			name: "Variation",
			s: &VariationRequest{
				Context:         ctx,
				Writers:         []io.WriteCloser{w},
				SkipClose:       skipClose,
				StartChannel:    startChannel,
				ContinueChannel: continueChannel,
				ErrorChannel:    errorChannel,
				ResultChannel:   resultChannel,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ctx, tt.s.GetContext())
			assert.Equal(t, w, tt.s.GetWriters()[0])
			assert.Equal(t, skipClose, tt.s.GetSkipClose())
			assert.Equal(t, startChannel, tt.s.GetStartChannel())
			assert.Equal(t, continueChannel, tt.s.GetContinueChannel())
			assert.Equal(t, errorChannel, tt.s.GetErrorChannel())
			assert.Equal(t, resultChannel, tt.s.GetResultChannel())
		})
	}
}
