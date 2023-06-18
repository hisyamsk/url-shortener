package middlewares

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/helpers"
)

type Middleware interface {
	ValidateRequest(schema any) fiber.Handler
}

type middleware struct {
	Validate *validator.Validate
}

func NewMiddleware(validate *validator.Validate) Middleware {
	return &middleware{
		Validate: validate,
	}
}

func (middleware *middleware) ValidateRequest(schema any) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.BodyParser(&schema)
		helpers.PanicIfError(err)

		err = middleware.Validate.Struct(schema)
		helpers.PanicIfError(err)

		return c.Next()
	}
}
