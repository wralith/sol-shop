package mw

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func EchoZerologValuesMiddleware(c echo.Context, v middleware.RequestLoggerValues) error {
	log.Info().
		Str("URI", v.URI).
		Int("status", v.Status).
		Msg("request")

	return nil
}
