package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/handlers"
	"github.com/hisyamsk/url-shortener/middlewares"
)

func V1Router(router fiber.Router, handlers *handlers.Handlers, middlewares middlewares.Middleware) {
	v1Router := router.Group("/v1")

	userRouter(v1Router, handlers.UserHandler, middlewares)
	urlRouter(v1Router, handlers.UrlHandler, middlewares)
}
