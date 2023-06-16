package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/models"
	"github.com/hisyamsk/url-shortener/repositories"
)

type UserService interface {
	FindAll() []*models.UserResponse
	Find(field string, val any) *models.UserResponse
	FindUrlsById(id int) []*models.UrlModel
	Create(user *models.UserModel) *models.UserResponse
	Update(user *models.UserModel) *models.UserResponse
	Delete(id int)
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) UserService {
	return &userService{repository}
}

func (service *userService) FindAll() []*models.UserResponse {
	usersEntities := service.repository.FindAll()
	var userResponse []*models.UserResponse

	for _, val := range usersEntities {
		userResponse = append(userResponse, helpers.UserEntityToResponse(val))
	}

	return userResponse
}
func (service *userService) Find(field string, val any) *models.UserResponse {
	userEntity, err := service.repository.Find(field, val)
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	return helpers.UserEntityToResponse(userEntity)
}
func (service *userService) FindUrlsById(id int) []*models.UrlModel {
	_, err := service.repository.Find("id", id)
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	urls := service.repository.FindUrlsById(id)
	var urlsResponse []*models.UrlModel
	for _, url := range urls {
		urlsResponse = append(urlsResponse, helpers.UrlEntityToResponse(url))
	}

	return urlsResponse
}
func (service *userService) Create(user *models.UserModel) *models.UserResponse {
	_, err := service.repository.Find("username", user.Username)
	if err != nil {
		hashedPassword := helpers.HashPassword(user.Password)
		userEntity := &entities.User{Username: user.Username, Password: hashedPassword}

		service.repository.Create(userEntity)
		return helpers.UserEntityToResponse(userEntity)
	}

	panic(fiber.NewError(fiber.StatusBadRequest, "username already exists"))
}
func (service *userService) Update(user *models.UserModel) *models.UserResponse {
	// check if user exists
	foundUser, err := service.repository.Find("id", user.ID)
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	// check if user updates the username
	if user.Username != "" {
		_, err = service.repository.Find("username", user.Username)
		if err == nil {
			panic(fiber.NewError(fiber.StatusBadRequest, "username already exists"))
		}
		foundUser.Username = user.Username
	}

	// check if user updates the password
	if user.Password != "" {
		err := helpers.ComparePassword(foundUser.Password, user.Password)
		if err == nil {
			panic(fiber.NewError(fiber.StatusBadRequest, "new password cannot be the same as the old one"))
		}
		hashedPassword := helpers.HashPassword(user.Password)
		foundUser.Password = hashedPassword
	}

	service.repository.Update(foundUser)
	return helpers.UserEntityToResponse(foundUser)
}
func (service *userService) Delete(id int) {
	_, err := service.repository.Find("id", id)
	if err != nil {
		panic(fiber.NewError(fiber.StatusNotFound, err.Error()))
	}

	service.repository.DeleteUrlsById(id)
	service.repository.Delete(id)
}
