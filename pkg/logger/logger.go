package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func InitLogger(logLevel string, jsonLogging bool) {
	if !jsonLogging {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	l := zerolog.InfoLevel

	if logLevel == "debug" {
		l = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(l)
}
