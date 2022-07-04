package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func Debug(msg string, v ...interface{}) {
	log.Debug().Msgf(msg, v)
}

func Info(msg string, v ...interface{}) {
	log.Info().Msgf(msg, v)
}

func Warn(msg string, v ...interface{}) {
	log.Warn().Msgf(msg, v)
}

func Error(msg string, v ...interface{}) {
	log.Error().Msgf(msg, v)
}

func Fatal(msg string, v ...interface{}) {
	log.Fatal().Msgf(msg, v)
}
