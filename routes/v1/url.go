package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/handlers"
	"github.com/hisyamsk/url-shortener/middlewares"
	"github.com/hisyamsk/url-shortener/models"
)

func urlRouter(router fiber.Router, handler handlers.UrlHandler, middlewares middlewares.Middleware) {
	urlRouter := router.Group("/urls")

	urlRouter.Get("/", handler.GetAll)
	urlRouter.Post("/", middlewares.ValidateRequest(&models.UrlCreateRequest{}), handler.Create)
	urlRouter.Get("/:id", handler.GetById)
	urlRouter.Patch("/:id", middlewares.ValidateRequest(&models.UrlUpdateRequest{}), handler.Update)
	urlRouter.Delete("/:id", handler.Delete)
}
