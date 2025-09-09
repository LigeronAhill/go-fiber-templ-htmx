package vacancy

import (
	"net/http"

	"github.com/LigeronAhill/go-fiber/pkg/tadaptor"
	"github.com/LigeronAhill/go-fiber/views/components"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type VacancyHandler struct {
	router       fiber.Router
	repository   *Repository
	customLogger *zerolog.Logger
}

func NewHandler(router fiber.Router, repository *Repository, customLogger *zerolog.Logger) {
	h := VacancyHandler{
		router,
		repository,
		customLogger,
	}
	vacancyGroup := h.router.Group("/vacancy")
	vacancyGroup.Post("/", h.createVacancy)
}

func (v *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	form := FromCtx(c)
	notification_status := components.NotificationSuccess
	status := http.StatusOK
	msg := "Вакансия успешно создана"
	if !form.IsValid() {
		errors := form.Validate()
		notification_status = components.NotificationFail
		status = http.StatusBadRequest
		msg = FormatErrors(errors)
	}
	if form.IsValid() {
		err := v.repository.AddVacancy(&form)
		if err != nil {
			v.customLogger.Error().Msg(err.Error())
			notification_status = components.NotificationFail
			msg = "Database internal error"
			status = http.StatusInternalServerError
		}
	}
	component := components.Notification(msg, notification_status)
	return tadaptor.Render(c, component, status)
}
