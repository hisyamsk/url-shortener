package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/models"
	"github.com/hisyamsk/url-shortener/services"
)

type UserHandler interface {
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	GetUrlsById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type userHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) UserHandler {
	return &userHandler{s}
}

func (handler *userHandler) GetAll(c *fiber.Ctx) error {
	res := handler.service.FindAll()

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, res)
}

func (handler *userHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	helpers.PanicIfError(err)

	res := handler.service.Find("id", id)

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, res)
}

func (handler *userHandler) GetUrlsById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	helpers.PanicIfError(err)

	res := handler.service.FindUrlsById(id)

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, res)
}

func (handler *userHandler) Create(c *fiber.Ctx) error {
	user := &models.UserModel{}
	err := c.BodyParser(&user)
	helpers.PanicIfError(err)

	res := handler.service.Create(user)

	return helpers.SendWebResponseSuccess(c, fiber.StatusCreated, res)
}

func (handler *userHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	helpers.PanicIfError(err)

	user := &models.UserModel{ID: uint(id)}
	err = c.BodyParser(&user)
	helpers.PanicIfError(err)

	res := handler.service.Update(user)

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, res)
}

func (handler *userHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	helpers.PanicIfError(err)

	handler.service.Delete(id)

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, "delete successful")
}
