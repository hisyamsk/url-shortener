package tests

import (
	"github.com/hisyamsk/url-shortener/database"
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/repositories"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Users []*entities.User
var Urls []*entities.Url
var UserRepo repositories.UserRepository

func TestInit() {
	DB = database.NewDB(database.DBTestName)
	UserRepo = repositories.NewUserRepository(DB)
}

func PopulateTables() {
	Users = []*entities.User{
		{Username: "hisyam", Password: "password1"},
		{Username: "setiadi", Password: "password2"},
		{Username: "kurniawan", Password: "password3"},
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
