package logging

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

func GetLogger() zerolog.Logger {
	zerolog.DurationFieldUnit = time.Second
	zerolog.TimeFieldFormat = time.RFC3339

	multi := zerolog.MultiLevelWriter(consoleWriter())
	logger := zerolog.New(multi).With().Timestamp().Logger()

	return logger
}

func consoleWriter() ConsoleWriter {
	var w ConsoleWriter

	writer := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05", NoColor: false}

	writer.PartsOrder = []string{
		zerolog.TimestampFieldName,
		zerolog.CallerFieldName,
		zerolog.LevelFieldName,
		zerolog.MessageFieldName,
	}

	w.ConsoleWriter = writer

	return w
}
