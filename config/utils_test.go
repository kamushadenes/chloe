package config

import (
	"github.com/kamushadenes/chloe/models"
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

func TestEnvOrDefaultWithOptions(t *testing.T) {
	// Given
	expected := "default"
	key := "NOT_SET_ENV_VAR"

	// When
	result := envOrDefaultWithOptions(key, expected, []string{"default", "other"})

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the default value")
}

func TestEnvOrDefaultImageSize(t *testing.T) {
	// Given
	expected := "256x256"
	key := "IMAGE_SIZE_ENV_VAR"

	// When
	result := envOrDefaultImageSize(key, expected)

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the default value")
}

func TestEnvOrDefaultCompletionModel(t *testing.T) {
	// Given
	expected := "gpt-3.5-turbo"
	key := "COMPLETION_MODEL_ENV_VAR"

	// When
	result := envOrDefaultCompletionModel(key, models.GPT35Turbo)

	// Then
	assert.Equal(t, expected, result.String(), "Value should be equal to the default value")
}

func TestEnvOrDefaultTranscriptionModel(t *testing.T) {
	// Given
	expected := "whisper-1"
	key := "TRANSCRIPTION_MODEL_ENV_VAR"

	// When
	result := envOrDefaultTranscriptionModel(key, models.Whisper1)

	// Then
	assert.Equal(t, expected, result.String(), "Value should be equal to the default value")
}

func TestEnvOrDefaultModerationModel(t *testing.T) {
	// Given
	expected := "text-moderation-stable"
	key := "MODERATION_MODEL_ENV_VAR"

	// When
	result := envOrDefaultModerationModel(key, models.TextModerationStable)

	// Then
	assert.Equal(t, expected, result.String(), "Value should be equal to the default value")
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

func TestEnvOrDefaultBool(t *testing.T) {
	// Given
	expected := true
	key := "BOOL_ENV_VAR"

	// When
	result := envOrDefaultBool(key, expected)

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the default value")
}

func TestEnvOrDefaultFloat64(t *testing.T) {
	// Given
	expected := 0.42
	key := "FLOAT_ENV_VAR"

	// When
	result := envOrDefaultFloat64(key, expected)

	// Then
	assert.Equal(t, expected, result, "Value should be equal to the default value")
}

func TestEnvOrDefaultIntInRange(t *testing.T) {
	// Given
	expected := 42
	key := "INT_IN_RANGE_ENV_VAR"

	// When
	result := envOrDefaultIntInRange(key, expected, 0, 100)

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
