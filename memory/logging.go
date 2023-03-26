package memory

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
	"time"
)

// Source: https://github.com/mpalmer/gorm-zerolog/blob/master/logger.go

type DBLogger struct {
	logMode logger.LogLevel
}

func (l DBLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.logMode = level
	return l
}

func (l DBLogger) Error(ctx context.Context, msg string, opts ...interface{}) {
	if l.logMode >= logger.Error {
		zerolog.Ctx(ctx).Error().Msg(fmt.Sprintf(msg, opts...))
	}
}

func (l DBLogger) Warn(ctx context.Context, msg string, opts ...interface{}) {
	if l.logMode >= logger.Warn {
		zerolog.Ctx(ctx).Warn().Msg(fmt.Sprintf(msg, opts...))
	}
}

func (l DBLogger) Info(ctx context.Context, msg string, opts ...interface{}) {
	if l.logMode >= logger.Info {
		zerolog.Ctx(ctx).Info().Msg(fmt.Sprintf(msg, opts...))
	}
}

func (l DBLogger) Trace(ctx context.Context, begin time.Time, f func() (string, int64), err error) {
	if l.logMode <= logger.Silent {
		return
	}

	zl := zerolog.Ctx(ctx)
	var event *zerolog.Event

	if err != nil {
		event = zl.Debug()
	} else {
		event = zl.Trace()
	}

	var durKey string

	switch zerolog.DurationFieldUnit {
	case time.Nanosecond:
		durKey = "elapsed_ns"
	case time.Microsecond:
		durKey = "elapsed_us"
	case time.Millisecond:
		durKey = "elapsed_ms"
	case time.Second:
		durKey = "elapsed"
	case time.Minute:
		durKey = "elapsed_min"
	case time.Hour:
		durKey = "elapsed_hr"
	default:
		zl.Error().Interface("zerolog.DurationFieldUnit", zerolog.DurationFieldUnit).Msg("gormzerolog encountered a mysterious, unknown value for DurationFieldUnit")
		durKey = "elapsed_"
	}

	event.Dur(durKey, time.Since(begin))

	sql, rows := f()
	if sql != "" {
		event.Str("sql", sql)
	}
	if rows > -1 {
		event.Int64("rows", rows)
	}

	event.Send()
}
