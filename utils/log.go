package utils

import "github.com/rs/zerolog/log"

func LogError(err error, msg string) {
	log.Error().Err(err).Msg(msg)
}

func LogData(data interface{}, msg string) {
	log.Info().Interface("data", data).Msg(msg)
}

func LogInfo(msg string) {
	log.Info().Msg(msg)
}

func LogFatal(err error, msg string) {
	log.Fatal().Err(err).Msg(msg)
}
