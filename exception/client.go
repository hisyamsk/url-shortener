package exception

import "github.com/gofiber/fiber/v2"

func RaiseIfNotFoundError(err error) {
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}
}

func RaiseIfDuplicateError(err error, message string) {
	if err == nil {
		panic(fiber.NewError(fiber.StatusBadRequest, message))
	}
}
