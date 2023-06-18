package helpers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/models"
)

func SendWebResponseSuccess(c *fiber.Ctx, code int, data any) error {
	return c.Status(code).JSON(&models.WebResponse{
		Success: true,
		Error:   nil,
		Data:    data,
	})
}
