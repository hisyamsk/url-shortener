package database

import (
	"fmt"
	"os"
	"time"

	"github.com/hisyamsk/url-shortener/config"
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(dbName string) *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: config.GormLogger,
	})
	helpers.PanicIfError(err)

	err = db.AutoMigrate(&entities.Url{}, &entities.User{})
	helpers.PanicIfError(err)

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	return db
}
