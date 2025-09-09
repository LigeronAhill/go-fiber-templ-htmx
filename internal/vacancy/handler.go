package vacancy

import (
	"github.com/LigeronAhill/go-fiber/pkg/logger/tadaptor"
	"github.com/LigeronAhill/go-fiber/views/components"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, customLogger *zerolog.Logger) {
	h := VacancyHandler{
		router,
		customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (v *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := FromCtx(c)
	status := components.NotificationSuccess
	msg := "Вакансия успешно создана"
	if !form.IsValid() {
		errors := form.Validate()
		status = components.NotificationFail
		msg = FormatErrors(errors)
	}
	component := components.Notification(msg, status)
	return tadaptor.Render(c, component)
}
