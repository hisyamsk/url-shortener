package tests

import (
	"github.com/hisyamsk/url-shortener/config"
	"github.com/hisyamsk/url-shortener/database"
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/repositories"
	"github.com/hisyamsk/url-shortener/services"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Users []*entities.User
var Urls []*entities.Url

var UserRepo repositories.UserRepository
var UrlRepo repositories.UrlRepository

var UserService services.UserService

func TestInit() {
	config.Init()
	DB = database.NewDB(database.DBTestName)
	UserRepo = repositories.NewUserRepository(DB)
	UrlRepo = repositories.NewUrlRepository(DB)
	UserService = services.NewUserService(UserRepo)
}

func PopulateTables() {
	pw := helpers.HashPassword("password123")
	Users = []*entities.User{
		{Username: "hisyam", Password: pw},
		{Username: "setiadi", Password: pw},
		{Username: "kurniawan", Password: pw},
	}
	err := DB.Create(&Users).Error
	helpers.PanicIfError(err)

	Urls = []*entities.Url{
		{Url: "url1", Redirect: "google.com", UserID: Users[0].ID},
		{Url: "url2", Redirect: "google.com", UserID: Users[0].ID},
		{Url: "url3", Redirect: "google.com", UserID: Users[1].ID},
	}
	err = DB.Create(&Urls).Error
	helpers.PanicIfError(err)
}

func DeleteRecords() {
	DB.Exec("TRUNCATE TABLE urls, users RESTART IDENTITY")
}
