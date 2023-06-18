package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/models"
	"github.com/hisyamsk/url-shortener/services"
)

type UrlHandler interface {
	GetAll(c *fiber.Ctx) error
	GetById(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type urlHandler struct {
	service services.UrlService
}

func NewUrlHandler(s services.UrlService) UrlHandler {
	return &urlHandler{s}
}

func (handler *urlHandler) GetAll(c *fiber.Ctx) error {
	res := handler.service.FindAll()

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, res)
}

func (handler *urlHandler) GetById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	helpers.PanicIfError(err)

	res := handler.service.Find("id", id)

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, res)
}

func (handler *urlHandler) Create(c *fiber.Ctx) error {
	var url *models.UrlModel
	err := c.BodyParser(&url)
	helpers.PanicIfError(err)

	res := handler.service.Create(url)

	return helpers.SendWebResponseSuccess(c, fiber.StatusCreated, res)
}

func (handler *urlHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	helpers.PanicIfError(err)

	url := &models.UrlModel{ID: uint(id)}
	err = c.BodyParser(&url)
	helpers.PanicIfError(err)

	res := handler.service.Update(url)

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, res)
}

func (handler *urlHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	helpers.PanicIfError(err)

	handler.service.Delete(id)

	return helpers.SendWebResponseSuccess(c, fiber.StatusOK, "delete successful")
}
