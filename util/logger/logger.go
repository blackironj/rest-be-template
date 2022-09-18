package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/blackironj/rest-be-template/env"
)

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "t"
	zerolog.LevelFieldName = "l"
	zerolog.MessageFieldName = "msg"

	logLevel := zerolog.DebugLevel
	switch env.LogLevel {
	case "info":
		logLevel = zerolog.InfoLevel
	case "warn":
		logLevel = zerolog.WarnLevel
	case "error":
		logLevel = zerolog.ErrorLevel
	case "fatal":
		logLevel = zerolog.FatalLevel
	case "panic":
		logLevel = zerolog.PanicLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	log.Logger = log.With().Caller().Timestamp().Logger()
}
