package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/handlers"
)

func UserRouter(router fiber.Router, handler handlers.UserHandler) {
	userRouter := router.Group("/users")

	userRouter.Post("/", handler.Create)
}
