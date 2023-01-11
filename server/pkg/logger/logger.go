package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger interface {
	Debug(message interface{}, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(err error, message string, args ...interface{})
	Fatal(err error, message string, args ...interface{})
}

type ZeroLogger struct {
	logger *zerolog.Logger
}

var _ Logger = (*ZeroLogger)(nil)

func New(level string) *ZeroLogger {
	var l zerolog.Level

	switch strings.ToLower(level) {
	case "error":
		l = zerolog.ErrorLevel
	case "warn":
		l = zerolog.WarnLevel
	case "info":
		l = zerolog.InfoLevel
	case "debug":
		l = zerolog.DebugLevel
	default:
		l = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(l)

	skipFrameCount := 3
	logger := zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()

	return &ZeroLogger{
		logger: &logger,
	}
}

func (l *ZeroLogger) Debug(message interface{}, args ...interface{}) {
	switch msg := message.(type) {
	case error:
		l.logDebug(msg.Error(), args...)
	case string:
		l.logDebug(msg, args...)
	default:
		l.logDebug(fmt.Sprintf("debug message %v has unknown type %v", message, msg), args...)
	}
}

func (l *ZeroLogger) logDebug(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Debug().Msg(message)
	} else {
		l.logger.Debug().Msgf(message, args...)
	}
}

func (l *ZeroLogger) Info(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Info().Msg(message)
	} else {
		l.logger.Info().Msgf(message, args...)
	}
}

func (l *ZeroLogger) Warn(message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Warn().Msg(message)
	} else {
		l.logger.Warn().Msgf(message, args...)
	}
}

func (l *ZeroLogger) Error(err error, message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Error().Err(err).Msg(message)
	} else {
		l.logger.Error().Err(err).Msgf(message, args...)
	}
}

func (l *ZeroLogger) Fatal(err error, message string, args ...interface{}) {
	if len(args) == 0 {
		l.logger.Fatal().Err(err).Msg(message)
	} else {
		l.logger.Fatal().Err(err).Msgf(message, args...)
	}

	os.Exit(1)
}
