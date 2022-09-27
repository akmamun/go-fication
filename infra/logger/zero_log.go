package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"strings"
)

var logger zerolog.Logger

func SetLogLevel(level string) {
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
	logger = zerolog.New(os.Stdout).With().Timestamp().CallerWithSkipFrameCount(zerolog.CallerSkipFrameCount + skipFrameCount).Logger()
}

func Info(message string, args ...interface{}) {
	if logger.GetLevel() == zerolog.InfoLevel {
		log.Info().Msgf(message, args...)
	}
}

func Debug(message string, args ...interface{}) {
	if logger.GetLevel() == zerolog.DebugLevel {
		log.Debug().Msgf(message, args...)
	}
}

func Warn(message string, args ...interface{}) {
	if logger.GetLevel() == zerolog.WarnLevel {
		log.Warn().Msgf(message, args...)
	}
}

func Error(message string, args ...interface{}) {
	if logger.GetLevel() == zerolog.ErrorLevel {
		log.Error().Msgf(message, args...)
	}
}

func Fatal(message string, args ...interface{}) {
	log.Fatal().Msgf(message, args)
	os.Exit(1)
}

func Log(message string, args ...interface{}) {
	if len(args) == 0 {
		log.Info().Msg(message)
	} else {
		log.Info().Msgf(message, args...)
	}
}
