// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

// Injectors from wire.go:

func InitializeServer(dbName string) *fiber.App {
	db := database.NewDB(dbName)
	urlRepository := repositories.NewUrlRepository(db)
	userRepository := repositories.NewUserRepository(db)
	urlService := services.NewUrlService(urlRepository, userRepository)
	mainHandler := handlers.NewMainHandler(urlService)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	urlHandler := handlers.NewUrlHandler(urlService)
	v1Handlers := &handlers.V1Handlers{
		UserHandler: userHandler,
		UrlHandler:  urlHandler,
	}
	apiVersionHandlers := &handlers.ApiVersionHandlers{
		MainHandler: mainHandler,
		V1Handlers:  v1Handlers,
	}
	validate := validator.New()
	middleware := middlewares.NewMiddleware(validate)
	fiberApp := app.NewApp(apiVersionHandlers, middleware)
	return fiberApp
}

// wire.go:

var userSet = wire.NewSet(repositories.NewUserRepository, services.NewUserService, handlers.NewUserHandler)

var urlSet = wire.NewSet(repositories.NewUrlRepository, services.NewUrlService, handlers.NewUrlHandler)
