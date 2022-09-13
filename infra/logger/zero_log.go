package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var logger zerolog.Logger

type LogLevel string

const (
	Infof  LogLevel = "Info"
	Warnf  LogLevel = "Warn"
	Debugf LogLevel = "Debug"
	Errorf LogLevel = "Error"
	Fatalf LogLevel = "Fatal"
)

func SetLogLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}

func Info(message string, args ...interface{}) {
	if logger.GetLevel() == zerolog.InfoLevel {
		logger.Info().Msgf(message, args...)
	}
}

func Debug(message string, args ...interface{}) {
	if logger.GetLevel() == zerolog.DebugLevel {
		logger.Debug().Msgf(message, args...)
	}
}

func Warn(message string, args ...interface{}) {
	if logger.GetLevel() == zerolog.WarnLevel {
		logger.Warn().Msgf(message, args...)
	}
}

func Error(message string, args ...interface{}) {
	if logger.GetLevel() == zerolog.ErrorLevel {
		logger.Error().Msgf(message, args...)
	}
}

func Fatal(message string, args ...interface{}) {
	logger.Fatal().Msgf(message, args)
	os.Exit(1)
}

func Log(message string, args ...interface{}) {
	if len(args) == 0 {
		logger.Info().Msg(message)
	} else {
		logger.Info().Msgf(message, args...)
	}
}
