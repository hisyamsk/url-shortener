package handlers

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/models"
	"github.com/hisyamsk/url-shortener/tests"
	"github.com/stretchr/testify/assert"
)

func TestUrlHandlerCreateSuccess(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	reqBody := &models.UrlCreateRequest{Url: "myCustomURL", Redirect: "https://google.com", UserID: 3}

	response, result := tests.GetResponse(fiber.MethodPost, "/v1/urls/", reqBody)
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data: map[string]interface{}{
			"id":       float64(4),
			"url":      reqBody.Url,
			"redirect": reqBody.Redirect,
			"userId":   float64(reqBody.UserID),
		},
	}

	assert.Equal(t, fiber.StatusCreated, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerCreateSuccessCustom(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	reqBody := &models.UrlCreateRequest{Redirect: "https://google.com", UserID: 3}

	response, result := tests.GetResponse(fiber.MethodPost, "/v1/urls/", reqBody)

	assert.Equal(t, fiber.StatusCreated, response.StatusCode)
	assert.Equal(t, true, result.Success)
	assert.Equal(t, nil, result.Error)
}

func TestUrlHandlerCreateFailedUrl(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	reqBody := &models.UrlCreateRequest{Url: "$/*@#$", Redirect: "https://google.com", UserID: 3}

	response, result := tests.GetResponse(fiber.MethodPost, "/v1/urls/", reqBody)
	expected := &models.WebResponse{
		Success: false,
		Error:   []interface{}{"error on field: Url, condition: alphanum"},
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerCreateFailedRedirect(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	reqBody := &models.UrlCreateRequest{Url: "myCustomURL", Redirect: "google.com", UserID: 3}

	response, result := tests.GetResponse(fiber.MethodPost, "/v1/urls/", reqBody)
	expected := &models.WebResponse{
		Success: false,
		Error:   []interface{}{"error on field: Redirect, condition: url"},
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerCreateFailedUserID(t *testing.T) {
	defer tests.DeleteRecords()
	reqBody := &models.UrlCreateRequest{Url: "myCustomURL", Redirect: "https://google.com", UserID: 1}

	response, result := tests.GetResponse(fiber.MethodPost, "/v1/urls/", reqBody)
	expected := &models.WebResponse{
		Success: false,
		Error:   "record not found",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerGetByIdSuccess(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodGet, "/v1/urls/1", nil)
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data: map[string]interface{}{
			"id":       float64(tests.Urls[0].ID),
			"url":      tests.Urls[0].Url,
			"redirect": tests.Urls[0].Redirect,
			"userId":   float64(tests.Urls[0].UserID),
		},
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerGetByIdFailed(t *testing.T) {
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodGet, "/v1/urls/1", nil)
	expected := &models.WebResponse{
		Success: false,
		Error:   "record not found",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerGetAll(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodGet, "/v1/urls", nil)
	expectedData := []interface{}{}
	for _, url := range tests.Urls {
		expectedData = append(expectedData, map[string]interface{}{
			"id":       float64(url.ID),
			"url":      url.Url,
			"redirect": url.Redirect,
			"userId":   float64(url.UserID),
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

func TestUrlHandlerUpdateSuccess(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	body := &models.UrlUpdateRequest{Url: "updatedUrl", UserID: 1}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/urls/1", body)
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data: map[string]interface{}{
			"id":       float64(tests.Urls[0].ID),
			"url":      "updatedUrl",
			"redirect": tests.Urls[0].Redirect,
			"userId":   float64(tests.Urls[0].UserID),
		},
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerUpdateFailedId(t *testing.T) {
	defer tests.DeleteRecords()
	body := &models.UrlUpdateRequest{Url: "myurl3", UserID: 1}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/urls/1", body)
	expected := &models.WebResponse{
		Success: false,
		Error:   "record not found",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerUpdateFailedUrlDuplicate(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	body := &models.UrlUpdateRequest{Url: "myurl3", UserID: 1}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/urls/1", body)
	expected := &models.WebResponse{
		Success: false,
		Error:   "url already exists",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerUpdateFailedValidation(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	body := &models.UrlUpdateRequest{Url: "myCustomURL", Redirect: "nohttps.com", UserID: 1}

	response, result := tests.GetResponse(fiber.MethodPatch, "/v1/urls/1", body)
	expected := &models.WebResponse{
		Success: false,
		Error:   []interface{}{"error on field: Redirect, condition: url"},
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusBadRequest, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerDeleteSuccess(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodDelete, "/v1/urls/1", nil)
	expected := &models.WebResponse{
		Success: true,
		Error:   nil,
		Data:    "delete successful",
	}

	assert.Equal(t, fiber.StatusOK, response.StatusCode)
	assert.Equal(t, expected, result)
}

func TestUrlHandlerDeleteFailedId(t *testing.T) {
	defer tests.DeleteRecords()

	response, result := tests.GetResponse(fiber.MethodDelete, "/v1/urls/1", nil)
	expected := &models.WebResponse{
		Success: false,
		Error:   "record not found",
		Data:    nil,
	}

	assert.Equal(t, fiber.StatusNotFound, response.StatusCode)
	assert.Equal(t, expected, result)
}
