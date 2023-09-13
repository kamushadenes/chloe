package config

import "time"

type TimeoutsConfig struct {
	Completion      time.Duration
	Transcription   time.Duration
	Moderation      time.Duration
	ImageGeneration time.Duration
	ImageEdit       time.Duration
	ImageVariation  time.Duration
	TTS             time.Duration
	SlownessWarning time.Duration
}

var Timeouts = &TimeoutsConfig{
	Completion:      envOrDefaultDuration("CHLOE_TIMEOUT_COMPLETION", 60*time.Second),
	Transcription:   envOrDefaultDuration("CHLOE_TIMEOUT_TRANSCRIPTION", 120*time.Second),
	Moderation:      envOrDefaultDuration("CHLOE_TIMEOUT_MODERATION", 60*time.Second),
	ImageGeneration: envOrDefaultDuration("CHLOE_TIMEOUT_IMAGE_GENERATION", 120*time.Second),
	ImageEdit:       envOrDefaultDuration("CHLOE_TIMEOUT_IMAGE_EDIT", 120*time.Second),
	ImageVariation:  envOrDefaultDuration("CHLOE_TIMEOUT_IMAGE_VARIATION", 120*time.Second),
	TTS:             envOrDefaultDuration("CHLOE_TIMEOUT_TTS", 60*time.Second),
	SlownessWarning: envOrDefaultDuration("CHLOE_TIMEOUT_SLOWNESS_WARNING", 5*time.Second),
}
