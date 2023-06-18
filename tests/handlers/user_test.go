package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/models"
	"github.com/hisyamsk/url-shortener/tests"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../../.env")
	helpers.PanicIfError(err)
	tests.TestInit()

	m.Run()

	defer tests.DeleteRecords()
}

func TestUserHandlerCreateSucess(t *testing.T) {
	defer tests.DeleteRecords()
	reqBody := &models.UserCreateRequest{Username: "hisyam", Password: "password123"}

	response, result := tests.GetResponse(fiber.MethodPost, "/v1/users", reqBody)
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data: map[string]interface{}{
			"id":       float64(1),
			"username": reqBody.Username,
		},
	}

	assert.Equal(t, fiber.StatusCreated, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerCreateFailedUsername(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	reqBody := &models.UserCreateRequest{Username: "hisyam", Password: "password123"}

	response, result := tests.GetResponse(fiber.MethodPost, "/v1/users", reqBody)
	expected := &models.WebResponse{
		Success: false,
		Error:   "username already exists",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerCreateFailedRequest(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	reqBody := &models.UserCreateRequest{Username: "hisyam"}

	response, result := tests.GetResponse(fiber.MethodPost, "/v1/users", reqBody)
	expected := &models.WebResponse{
		Success: false,
		Error: []interface{}{
			"error on field: Password, condition: required",
		},
		Data: nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerGetByIdSuccess(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodGet, "/v1/users/1", nil)
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data: map[string]interface{}{
			"id":       float64(tests.Users[0].ID),
			"username": tests.Users[0].Username,
		},
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerGetByIdFailed(t *testing.T) {
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodGet, "/v1/users/1", nil)
	expected := &models.WebResponse{
		Success: false,
		Error:   "record not found",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerGetAll(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodGet, "/v1/users", nil)
	expectedData := []interface{}{}
	for _, item := range tests.Users {
		expectedData = append(expectedData, map[string]interface{}{
			"id":       float64(item.ID),
			"username": item.Username,
		})
	}
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data:    expectedData,
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerUpdateSuccess(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	req := &models.UserUpdateRequest{Username: "hisyam_2"}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/users/1", req)
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data: map[string]interface{}{
			"id":       float64(1),
			"username": "hisyam_2",
		},
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerUpdateFailedID(t *testing.T) {
	defer tests.DeleteRecords()
	req := &models.UserUpdateRequest{Username: "hisyam_2"}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/users/1", req)
	expected := &models.WebResponse{
		Success: false,
		Error:   "record not found",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerUpdateFailedPassword(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	req := &models.UserUpdateRequest{Password: "password123"}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/users/1", req)
	expected := &models.WebResponse{
		Success: false,
		Error:   "new password cannot be the same as the old one",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerUpdateFailedUsername(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	req := &models.UserUpdateRequest{Username: "setiadi"}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/users/1", req)
	expected := &models.WebResponse{
		Success: false,
		Error:   "username already exists",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerUpdateFailedRequest(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	req := &models.UserUpdateRequest{Username: "a"}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/users/1", req)
	expected := &models.WebResponse{
		Success: false,
		Error: []interface{}{
			"error on field: Username, condition: min",
		},
		Data: nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerGetUrlsById(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodGet, "/v1/users/1/urls", nil)
	expectedData := []interface{}{}
	for _, url := range tests.Urls {
		if url.UserID == 1 {
			expectedData = append(expectedData, map[string]interface{}{
				"id":       float64(url.ID),
				"url":      url.Url,
				"redirect": url.Redirect,
				"userId":   float64(url.UserID),
			})
		}
	}
	expected := &models.WebResponse{Success: true, Error: nil, Data: expectedData}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerGetUrlsByIdFailed(t *testing.T) {
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodGet, "/v1/users/1/urls", nil)
	expected := &models.WebResponse{
		Success: false,
		Error:   "record not found",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerDeleteSucess(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodDelete, "/v1/users/1", nil)
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data:    "delete successful",
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUserHandlerDeleteFailed(t *testing.T) {
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodDelete, "/v1/users/1", nil)
	expected := &models.WebResponse{
		Success: false,
		Error:   "record not found",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}
