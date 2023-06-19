package main

import (
	"os"

	"github.com/hisyamsk/url-shortener/app/database"
	"github.com/hisyamsk/url-shortener/app/server"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	helpers.PanicIfError(err)

	server := server.InitializeServer(database.DBName)
	err = server.Listen(os.Getenv("APP_PORT"))
	helpers.PanicIfError(err)
}
