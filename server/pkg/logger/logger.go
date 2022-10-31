package logger

import (
	"os"
	"sol-shop-server/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger(c *config.LoggerConfig) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	if c.Pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if c.Level == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}
