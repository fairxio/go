package log

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func Debug(msg string, v ...interface{}) {
	if v != nil {
		log.Debug().Msgf(msg, v)
		return
	}
	log.Debug().Msg(msg)
}

func Info(msg string, v ...interface{}) {
	if v != nil {
		log.Info().Msgf(msg, v)
		return
	}
	log.Info().Msg(msg)
}

func Warn(msg string, v ...interface{}) {
	if v != nil {
		log.Warn().Msgf(msg, v)
		return
	}
	log.Warn().Msg(msg)
}

func Error(msg string, v ...interface{}) {
	if v != nil {
		log.Error().Msgf(msg, v)
		return
	}
	log.Error().Msg(msg)
}

func Fatal(msg string, v ...interface{}) {
	if v != nil {
		log.Fatal().Msgf(msg, v)
		return
	}
	log.Fatal().Msg(msg)
}
