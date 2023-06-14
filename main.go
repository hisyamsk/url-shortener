package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/hisyamsk/url-shortener/database"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helpers.PanicIfError(err)
	database.NewDB(database.DBName)

	app := fiber.New()
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, World",
		})
	})

	err = app.Listen(":8000")
	helpers.PanicIfError(err)
}
