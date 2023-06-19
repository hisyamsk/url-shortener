package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/services"
)

type MainHandler interface {
	FindUrl(c *fiber.Ctx) error
}

type handler struct {
	UrlService services.UrlService
}

func NewMainHandler(urlService services.UrlService) MainHandler {
	return &handler{
		UrlService: urlService,
	}
}

func (handler *handler) FindUrl(c *fiber.Ctx) error {
	url := c.Params("url")

	foundUrl := handler.UrlService.Find("url", url)
	return c.Redirect(foundUrl.Redirect, fiber.StatusFound)
}
