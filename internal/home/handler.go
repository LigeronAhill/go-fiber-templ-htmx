package home

import (
	"net/http"

	"github.com/LigeronAhill/go-fiber/pkg/tadaptor"
	"github.com/LigeronAhill/go-fiber/views"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router fiber.Router
	log    *zerolog.Logger
}

func NewHandler(router fiber.Router, logger *zerolog.Logger) {
	h := &HomeHandler{
		router: router,
		log:    logger,
	}
	h.router.Get("/", h.home)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Main()
	return tadaptor.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "AMAZING ERROR")
}
