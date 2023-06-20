package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hisyamsk/url-shortener/config"
	"github.com/hisyamsk/url-shortener/handlers"
	"github.com/hisyamsk/url-shortener/middlewares"
	"github.com/hisyamsk/url-shortener/routes"
)

func NewApp(handlers *handlers.ApiVersionHandlers, middlewares middlewares.Middleware) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: config.FiberErrorHandler,
	})

	// middlewares
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	routes.Router(app, handlers, middlewares)

	return app
}
