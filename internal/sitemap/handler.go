package sitemap

import (
	"bytes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sabloger/sitemap-generator/smg"
)

type SitemapHandler struct {
	router fiber.Router
}

func NewHandler(router fiber.Router) {
	h := SitemapHandler{
		router,
	}
	h.router.Get("/sitemap.xml", h.sitemap)
}

func (v *SitemapHandler) sitemap(c *fiber.Ctx) error {
	now := time.Now().UTC()
	sm := smg.NewSitemap(true)
	sm.SetHostname("https://example.com")
	sm.SetLastMod(&now)
	sm.SetCompress(false)
	err := sm.Add(&smg.SitemapLoc{
		Loc:        "/",
		LastMod:    &now,
		ChangeFreq: smg.Daily,
		Priority:   0.8,
	})
	if err != nil {
		return err
	}
	err = sm.Add(&smg.SitemapLoc{
		Loc:        "/login",
		LastMod:    &now,
		ChangeFreq: smg.Weekly,
		Priority:   0.6,
	})
	if err != nil {
		return err
	}
	sm.Finalize()

	var buf bytes.Buffer

	_, err = sm.WriteTo(&buf)
	if err != nil {
		return err
	}
	c.Set("Content-Type", "application/xml")
	return c.Send(buf.Bytes())
}
