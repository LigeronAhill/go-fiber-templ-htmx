package main

import (
	"github.com/gofiber/contrib/fiberzerolog"

	"github.com/LigeronAhill/go-fiber/config"
	"github.com/LigeronAhill/go-fiber/internal/home"
	"github.com/LigeronAhill/go-fiber/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func main() {
	config.Init()
	logCfg := config.NewLogConfig()
	log := logger.CreateLogger(logCfg)

	log.Info().Int("Setting log level", logCfg.Level).Send()

	dbCfg := config.NewDatabaseConfig()
	log.Debug().Str("Database URL set to", dbCfg.URL).Send()

	loggerConfig := logger.New(logCfg)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(fiberzerolog.New(loggerConfig))
	app.Use(recover.New())

	home.NewHandler(app, log)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal().AnErr("start server error", err)
	}
}
