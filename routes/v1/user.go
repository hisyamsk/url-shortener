package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/handlers"
	"github.com/hisyamsk/url-shortener/middlewares"
	"github.com/hisyamsk/url-shortener/models"
)

func userRouter(router fiber.Router, handler handlers.UserHandler, middlewares middlewares.Middleware) {
	userRouter := router.Group("/users")

	userRouter.Get("/", handler.GetAll)
	userRouter.Post("/", middlewares.ValidateRequest(&models.UserCreateRequest{}), handler.Create)
	userRouter.Get("/:id", handler.GetById)
	userRouter.Patch("/:id", middlewares.ValidateRequest(&models.UserUpdateRequest{}), handler.Update)
	userRouter.Delete("/:id", handler.Delete)
	userRouter.Get("/:id/urls", handler.GetUrlsById)
}
