package home

import (
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
	api := h.router.Group("/api")
	api.Get("/", h.home)
	api.Get("/error", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "AMAZING ERROR")
}
