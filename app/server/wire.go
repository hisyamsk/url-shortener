//go:build wireinject
// +build wireinject

package server

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/hisyamsk/url-shortener/app"
	"github.com/hisyamsk/url-shortener/app/database"
	"github.com/hisyamsk/url-shortener/handlers"
	"github.com/hisyamsk/url-shortener/middlewares"
	"github.com/hisyamsk/url-shortener/repositories"
	"github.com/hisyamsk/url-shortener/services"
)

var userSet = wire.NewSet(
	repositories.NewUserRepository,
	services.NewUserService,
	handlers.NewUserHandler,
)

var urlSet = wire.NewSet(
	repositories.NewUrlRepository,
	services.NewUrlService,
	handlers.NewUrlHandler,
)

func InitializeServer(dbName string) *fiber.App {
	wire.Build(
		database.NewDB,
		validator.New,
		userSet,
		urlSet,
		wire.Struct(new(handlers.V1Handlers), "*"),
		wire.Struct(new(handlers.ApiVersionHandlers), "*"),
		middlewares.NewMiddleware,
		app.NewApp,
	)

	return nil
}
