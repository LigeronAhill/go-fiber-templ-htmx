package main

import (
	"os"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/storage/postgres/v3"

	"github.com/LigeronAhill/go-fiber/config"
	"github.com/LigeronAhill/go-fiber/internal/home"
	"github.com/LigeronAhill/go-fiber/internal/sitemap"
	"github.com/LigeronAhill/go-fiber/internal/vacancy"
	"github.com/LigeronAhill/go-fiber/pkg/database"
	"github.com/LigeronAhill/go-fiber/pkg/logger"
	"github.com/LigeronAhill/go-fiber/pkg/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func main() {
	config.Init()
	logCfg := config.NewLogConfig()
	log := logger.CreateLogger(logCfg)

	log.Info().Int("Setting log level", logCfg.Level).Send()

	dbCfg := config.NewDatabaseConfig()
	pool, err := database.CreateDBPool(dbCfg)
	if err != nil {
		log.Error().Msg(err.Error())
		os.Exit(1)
	}
	defer pool.Close()
	log.Info().Msg("Database connected")
	vacancyRepo := vacancy.NewRepo(pool, log)
	storage := postgres.New(postgres.Config{
		DB:         pool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})

	store := session.New(session.Config{
		Storage: storage,
	})

	loggerConfig := logger.New(logCfg)

	app := fiber.New()
	app.Use(requestid.New())
	app.Use(fiberzerolog.New(loggerConfig))
	app.Use(recover.New())

	app.Use(middleware.AuthMiddleware(store))

	app.Static("/public", "./public")

	home.NewHandler(app, store, vacancyRepo, log)
	vacancy.NewHandler(app, vacancyRepo, log)
	sitemap.NewHandler(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatal().AnErr("start server error", err)
	}
}
