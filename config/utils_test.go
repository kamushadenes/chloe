package config

import (
	"os"
	"testing"
	"time"

	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"github.com/stretchr/testify/assert"
)

func TestEnvOrDefault(t *testing.T) {
	// Given
	expected := "default"
	key := "NOT_SET_ENV_VAR"

	// When
	result := envOrDefault(key, expected)

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the default value")
}

func TestEnvOrDefaultGCPTTSEncoding(t *testing.T) {
	// Given
	expected := texttospeechpb.AudioEncoding_LINEAR16
	key := "GCP_TTS_ENCODING"

	// When
	result := envOrDefaultGCPTTSEncoding(key, expected)

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the default value")
}

func TestEnvOrDefaultDuration(t *testing.T) {
	// Given
	expected := 10 * time.Second
	key := "DURATION_ENV_VAR"

	// When
	result := envOrDefaultDuration(key, expected)

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the default value")
}

func TestEnvOrDefaultInt(t *testing.T) {
	// Given
	expected := 42
	key := "INT_ENV_VAR"

	// When
	result := envOrDefaultInt(key, expected)

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the default value")
}

func TestMustEnv(t *testing.T) {
	// Given
	expected := "some value"

	_ = os.Setenv("MUST_SET_ENV_VAR", expected)

	key := "MUST_SET_ENV_VAR"

	// When
	result := mustEnv(key)

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the expected value")
	assert.Panicsf(t, func() { mustEnv("SOME_OTHER_ENV_VARIABLE") }, "should panic when key not found")
}
