package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTickerOrDefault(t *testing.T) {
	testCases := []struct {
		name          string
		inputDuration time.Duration
		defDuration   time.Duration
		expected      time.Duration
	}{
		{
			name:          "Valid duration and default",
			inputDuration: 2 * time.Second,
			defDuration:   5 * time.Second,
			expected:      2 * time.Second,
		},
		{
			name:          "Zero duration with valid default",
			inputDuration: 0,
			defDuration:   5 * time.Second,
			expected:      5 * time.Second,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ticker := TickerOrDefault(tc.inputDuration, tc.defDuration)

			now := time.Now()
			<-ticker.C
			assert.True(t, time.Since(now) >= tc.expected)

			ticker.Stop()
		})
	}
}
