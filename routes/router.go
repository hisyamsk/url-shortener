package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/handlers"
	v1 "github.com/hisyamsk/url-shortener/routes/v1"
)

func Router(app *fiber.App, handlers *handlers.ApiVersionHandlers) {
	v1.V1Router(app, handlers.V1Handlers)
	app.Route("/v2", func(router fiber.Router) {
		router.Get("/users", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"success": true,
				"error":   nil,
				"data":    nil,
			})
		})
	})
}
