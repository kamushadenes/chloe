package config

import (
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
	"fmt"
	"github.com/kamushadenes/chloe/errors"
	"github.com/kamushadenes/chloe/models"
	"github.com/kamushadenes/chloe/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func envOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}

func envOrDefaultWithOptions(key, defaultValue string, options []string) string {
	value := envOrDefault(key, defaultValue)
	if !utils.StringInSlice(value, options) {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", key),
			fmt.Errorf("value: %s", value),
			fmt.Errorf("options: %s", strings.Join(options, ", ")),
		))
	}

	return value
}

func envOrDefaultImageSize(key, defaultValue string) string {
	return envOrDefaultWithOptions(key, defaultValue,
		[]string{
			"256x256",
			"512x512",
			"1024x1024",
		})
}

func envOrDefaultCompletionModel(key string, defaultValue *models.Model) *models.Model {
	return models.GetModel(envOrDefaultWithOptions(key, defaultValue.String(),
		models.ModelsToString(
			models.GPT35Turbo,
			models.GPT35Turbo0301,
			models.GPT4,
			models.GPT40314,
			models.GPT432K,
			models.GPT432K0314,
		)))
}

func envOrDefaultTranscriptionModel(key string, defaultValue *models.Model) *models.Model {
	return models.GetModel(envOrDefaultWithOptions(key, defaultValue.String(),
		models.ModelsToString(
			models.Whisper1,
		)))
}

func envOrDefaultModerationModel(key string, defaultValue *models.Model) *models.Model {
	return models.GetModel(envOrDefaultWithOptions(key, defaultValue.String(),
		models.ModelsToString(
			models.TextModerationStable,
			models.TextModerationLatest,
		)))
}

func envOrDefaultGCPTTSEncoding(key string, defaultValue texttospeechpb.AudioEncoding) texttospeechpb.AudioEncoding {
	if value := os.Getenv(key); value != "" {
		switch value {
		case "LINEAR16":
			return texttospeechpb.AudioEncoding_LINEAR16
		case "MP3":
			return texttospeechpb.AudioEncoding_MP3
		case "OGG_OPUS":
			return texttospeechpb.AudioEncoding_OGG_OPUS
		case "MULAW":
			return texttospeechpb.AudioEncoding_MULAW
		case "ALAW":
			return texttospeechpb.AudioEncoding_ALAW
		default:
			panic(errors.Wrap(
				errors.ErrInvalidEnv,
				fmt.Errorf("env: %s", key),
				fmt.Errorf("value: %s", value),
			))
		}
	}

	return defaultValue
}

func envOrDefaultDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		dur, err := time.ParseDuration(value)
		if err != nil {
			panic(errors.Wrap(
				errors.ErrInvalidEnv,
				err,
				fmt.Errorf("env: %s", key),
				fmt.Errorf("value: %s", value),
			))
		}
		return dur
	}

	return defaultValue
}

func envOrDefaultBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		b, err := strconv.ParseBool(value)
		if err != nil {
			panic(errors.Wrap(
				errors.ErrInvalidEnv,
				err,
				fmt.Errorf("env: %s", key),
				fmt.Errorf("value: %s", value),
			))
		}
		return b
	}

	return defaultValue
}

func envOrDefaultInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		i, err := strconv.Atoi(value)
		if err != nil {
			panic(errors.Wrap(
				fmt.Errorf("env: %s", key),
				fmt.Errorf("value: %s", value),
			))
		}
		return i
	}

	return defaultValue
}

func envOrDefaultIntInRange(key string, defaultValue, min, max int) int {
	value := envOrDefaultInt(key, defaultValue)
	if value < min || value > max {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", key),
			fmt.Errorf("value: %d", value),
			fmt.Errorf("min: %d", min),
			fmt.Errorf("max %d", max),
		))
	}

	return value
}

func envOrDefaultFloat64(key string, defaultValue float64) float64 {
	if value := os.Getenv(key); value != "" {
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			panic(errors.Wrap(
				errors.ErrInvalidEnv,
				fmt.Errorf("env: %s", key),
				fmt.Errorf("value: %s", value),
				err,
			))
		}
		return f
	}

	return defaultValue
}

func envOrDefaultFloat64InRange(key string, defaultValue, min, max float64) float64 {
	value := envOrDefaultFloat64(key, defaultValue)
	if value < min || value > max {
		panic(errors.Wrap(
			errors.ErrInvalidEnv,
			fmt.Errorf("env: %s", key),
			fmt.Errorf("value: %f", value),
			fmt.Errorf("min: %f", min),
			fmt.Errorf("max %f", max),
		))
	}

	return value
}

func mustEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(errors.Wrap(
			errors.ErrMissingEnv,
			fmt.Errorf("environment variable %s is not set", key),
		))
	}

	return value
}
