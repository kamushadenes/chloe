package logging

import (
	"context"
	"fmt"
	"github.com/kamushadenes/chloe/flags"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"os"
	"time"
)

func GetLogger() zerolog.Logger {
	zerolog.DurationFieldUnit = time.Second
	zerolog.TimeFieldFormat = time.RFC3339

	wr := diode.NewWriter(consoleWriter(), 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Dropped %d messages", missed)
	})

	multi := zerolog.MultiLevelWriter(wr)
	logger := zerolog.New(multi).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if flags.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	return logger
}

func FromContext(ctx context.Context) *zerolog.Logger {
	l := zerolog.Ctx(ctx)
	if l != nil {
		return l
	}

	logger := GetLogger()

	return &logger
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

	writer.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("| %-60s|", i)
	}

	w.ConsoleWriter = writer

	return w
}
