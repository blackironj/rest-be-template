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
}

func logEvent(level zerolog.Level, logFields ...map[string]interface{}) *zerolog.Event {
	var e *zerolog.Event

	switch level {
	case zerolog.DebugLevel:
		e = log.Debug()
	case zerolog.InfoLevel:
		e = log.Info()
	case zerolog.WarnLevel:
		e = log.Warn()
	case zerolog.ErrorLevel:
		e = log.Error()
	case zerolog.FatalLevel:
		e = log.Fatal()
	case zerolog.PanicLevel:
		e = log.Panic()
	}

	if len(logFields) > 0 {
		e.Fields(logFields)
	}
	return e
}

func Debug(msg string, logFields ...map[string]interface{}) {
	logEvent(zerolog.DebugLevel, logFields...).Msg(msg)
}

func Info(msg string, logFields ...map[string]interface{}) {
	logEvent(zerolog.InfoLevel, logFields...).Msg(msg)
}

func Warn(msg string, logFields ...map[string]interface{}) {
	logEvent(zerolog.WarnLevel, logFields...).Msg(msg)
}

func Error(msg string, logFields ...map[string]interface{}) {
	logEvent(zerolog.ErrorLevel, logFields...).Msg(msg)
}

func Fatal(msg string, logFields ...map[string]interface{}) {
	logEvent(zerolog.FatalLevel, logFields...).Msg(msg)
}

func Panic(msg string, logFields ...map[string]interface{}) {
	logEvent(zerolog.PanicLevel, logFields...).Msg(msg)
}
