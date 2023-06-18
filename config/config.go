package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/models"
	"gorm.io/gorm/logger"
)

var GormLogger logger.Interface = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      false,       // Don't include params in the SQL log
		Colorful:                  true,        // Disable color
	},
)

var FiberErrorHandler fiber.ErrorHandler = func(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	//  validation error
	if errors, ok := err.(validator.ValidationErrors); ok {
		errorMessages := []string{}
		for _, e := range errors {
			errorMessage := fmt.Sprintf("error on field: %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		return c.Status(fiber.StatusBadRequest).JSON(&models.WebResponse{
			Success: false,
			Error:   errorMessages,
			Data:    nil,
		})
	}

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	return c.Status(code).JSON(&models.WebResponse{
		Success: false,
		Error:   err.Error(),
		Data:    nil,
	})
}
