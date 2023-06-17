package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/handlers"
)

func V1Router(router fiber.Router, handlers *handlers.Handlers) {
	v1Router := router.Group("/v1")

	UserRouter(v1Router, handlers.UserHandler)
}
