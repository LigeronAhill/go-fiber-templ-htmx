package home

import (
	"net/http"

	"github.com/LigeronAhill/go-fiber/internal/vacancy"
	"github.com/LigeronAhill/go-fiber/pkg/tadaptor"
	"github.com/LigeronAhill/go-fiber/views"
	"github.com/LigeronAhill/go-fiber/views/components"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router fiber.Router
	store  *session.Store
	repo   *vacancy.Repository
	log    *zerolog.Logger
}

func NewHandler(router fiber.Router, store *session.Store, repo *vacancy.Repository, logger *zerolog.Logger) {
	h := &HomeHandler{
		router: router,
		store:  store,
		repo:   repo,
		log:    logger,
	}
	h.router.Get("/", h.home)
	h.router.Get("/login", h.login)
	h.router.Get("/logout", h.logout)
	h.router.Post("/login", h.loginForm)
	h.router.Get("/404", h.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	totalPages, err := h.repo.TotalPages()
	if err != nil {
		h.log.Error().Msg(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}
	page := max(min(c.QueryInt("page", 1), totalPages), 1)
	if page == 0 {
		page = 1
	}
	vacancies, err := h.repo.GetAll(page)
	if err != nil {
		h.log.Error().Msg(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}
	component := views.Main(vacancies, page, totalPages)
	return tadaptor.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return fiber.NewError(fiber.StatusBadRequest, "AMAZING ERROR")
}
func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tadaptor.Render(c, component, http.StatusOK)
}
func (h *HomeHandler) loginForm(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if email == "master@mail.com" && password == "my_pass" {
		sess, err := h.store.Get(c)
		if err != nil {
			h.log.Error().Msg(err.Error())
		}
		sess.Set("email", email)
		if err = sess.Save(); err != nil {
			h.log.Error().Msg(err.Error())
		}
		c.Response().Header.Add("Hx-Redirect", "/")
		return c.Redirect("/", http.StatusOK)
	}
	notification_status := components.NotificationFail
	status := http.StatusBadRequest
	msg := "Неверный логин или пароль"
	component := components.Notification(msg, notification_status)
	return tadaptor.Render(c, component, status)
}
func (h *HomeHandler) logout(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil {
		h.log.Error().Msg(err.Error())
	}
	sess.Delete("email")
	if err = sess.Save(); err != nil {
		h.log.Error().Msg(err.Error())
	}
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}
