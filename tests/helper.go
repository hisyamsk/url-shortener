package tests

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/app"
	"github.com/hisyamsk/url-shortener/app/database"
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/handlers"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/middlewares"
	"github.com/hisyamsk/url-shortener/models"
	"github.com/hisyamsk/url-shortener/repositories"
	"github.com/hisyamsk/url-shortener/routes"
	"github.com/hisyamsk/url-shortener/services"
	"gorm.io/gorm"
)

var DB *gorm.DB
var Users []*entities.User
var Urls []*entities.Url

var UserRepo repositories.UserRepository
var UrlRepo repositories.UrlRepository

var UserService services.UserService
var UrlService services.UrlService

var AppTest *fiber.App

func TestInit() {
	DB = database.NewDB(database.DBTestName)

	UserRepo = repositories.NewUserRepository(DB)
	UrlRepo = repositories.NewUrlRepository(DB)

	UserService = services.NewUserService(UserRepo)
	UrlService = services.NewUrlService(UrlRepo, UserRepo)

	UserHandler := handlers.NewUserHandler(UserService)
	UrlHandler := handlers.NewUrlHandler(UrlService)

	v1Handlers := &handlers.V1Handlers{
		UserHandler: UserHandler,
		UrlHandler:  UrlHandler,
	}
	rootHandlers := &handlers.ApiVersionHandlers{
		V1Handlers: v1Handlers,
	}

	validator := validator.New()
	Middlewares := middlewares.NewMiddleware(validator)
	AppTest = app.NewApp(rootHandlers, Middlewares)
	routes.Router(AppTest, rootHandlers, Middlewares)
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
		{Url: "myurl3", Redirect: "google.com", UserID: Users[1].ID},
	}
	err = DB.Create(&Urls).Error
	helpers.PanicIfError(err)
}

func GetResponse(method, target string, body any) (*http.Response, *models.WebResponse) {
	jsonBody, _ := json.Marshal(&body)
	reqBody := strings.NewReader(string(jsonBody))
	req := httptest.NewRequest(method, target, reqBody)
	req.Header.Set("Content-Type", "application/json")

	response, _ := AppTest.Test(req, -1)
	respBody, _ := io.ReadAll(response.Body)
	var result *models.WebResponse
	err := json.Unmarshal(respBody, &result)
	helpers.PanicIfError(err)

	return response, result
}

func DeleteRecords() {
	DB.Exec("TRUNCATE TABLE urls, users RESTART IDENTITY")
}
