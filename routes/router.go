package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/handlers"
	"github.com/hisyamsk/url-shortener/middlewares"
	v1 "github.com/hisyamsk/url-shortener/routes/v1"
)

func Router(app *fiber.App, handlers *handlers.ApiVersionHandlers, middlewares middlewares.Middleware) {
	app.Get("/:url", handlers.MainHandler.FindUrl)
	app.Static("/docs", "./static")

	v1.V1Router(app, handlers.V1Handlers, middlewares)
}
