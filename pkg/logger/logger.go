package logger

import (
	"os"

	"github.com/LigeronAhill/go-fiber/config"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func New(c *config.LogConfig) fiberzerolog.Config {
	zerolog.SetGlobalLevel(zerolog.Level(zerolog.Level(c.Level)))
	loggerConfig := fiberzerolog.ConfigDefault
	loggerConfig.GetLogger = func(f *fiber.Ctx) zerolog.Logger {
		id := f.Locals("requestid").(string)
		l := CreateLogger(c)
		return l.With().Str("requestID", id).Logger()
	}
	return loggerConfig
}

func CreateLogger(c *config.LogConfig) *zerolog.Logger {
	var l zerolog.Logger
	if c.Format == "json" {
		l = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		cw := zerolog.ConsoleWriter{Out: os.Stdout}
		l = zerolog.New(cw).With().Timestamp().Logger()
	}
	return &l
}
