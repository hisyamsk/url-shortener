package services

import (
	"testing"

	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/models"
	"github.com/hisyamsk/url-shortener/tests"
	"github.com/stretchr/testify/assert"
)

func TestUrlServiceCreate(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	url := &models.UrlModel{Redirect: "google.com", UserID: 1}

	result := tests.UrlService.Create(url)
	expected := &models.UrlModel{ID: 4, Redirect: "google.com", UserID: 1}

	assert.Equal(t, expected.ID, result.ID)
	assert.Equal(t, expected.Redirect, result.Redirect)
	assert.Equal(t, expected.UserID, result.UserID)
}

func TestUrlServiceCreateCustom(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	url := &models.UrlModel{Url: "customUrl", Redirect: "google.com", UserID: 1}

	result := tests.UrlService.Create(url)
	expected := &models.UrlModel{ID: 4, Url: "customUrl", Redirect: "google.com", UserID: 1}

	assert.Equal(t, expected, result)
}

func TestUrlServiceCreateFailedUser(t *testing.T) {
	defer tests.DeleteRecords()
	url := &models.UrlModel{Url: "customUrl", Redirect: "google.com", UserID: 1}

	assert.Panics(t, func() {
		tests.UrlService.Create(url)
	})
}

func TestUrlServiceCreateFailedUrl(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	url := &models.UrlModel{Url: "url1", Redirect: "google.com", UserID: 2}

	assert.Panics(t, func() {
		tests.UrlService.Create(url)
	})
}

func TestUrlServiceFind(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	result := tests.UrlService.Find("id", id)
	expected := helpers.UrlEntityToResponse(tests.Urls[0])

	assert.Equal(t, expected, result)
}

func TestUrlServiceFindFailed(t *testing.T) {
	defer tests.DeleteRecords()
	id := 1

	assert.Panics(t, func() {
		tests.UrlService.Find("id", id)
	})
}

func TestUrlServiceFindAll(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	var expected []*models.UrlModel

	result := tests.UrlService.FindAll()
	for _, val := range tests.Urls {
		expected = append(expected, helpers.UrlEntityToResponse(val))
	}

	assert.Equal(t, expected, result)
}

func TestUrlServiceUpdate(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	url := &models.UrlModel{ID: 1, Url: "UpdatedURL", UserID: 1}

	result := tests.UrlService.Update(url)
	expected := &models.UrlModel{ID: 1, Url: "UpdatedURL", Redirect: "google.com", UserID: 1}

	assert.Equal(t, expected, result)
}

func TestUrlServiceUpdateFailed(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	url := &models.UrlModel{ID: 1, Url: "url2", UserID: 1}

	assert.Panics(t, func() {
		tests.UrlService.Update(url)
	})
}

func TestUrlServiceDelete(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	tests.UserService.Delete(id)
	_, err := tests.UserRepo.Find("id", id)

	assert.NotNil(t, err)
}

func TestUrlServiceDeleteFailed(t *testing.T) {
	defer tests.DeleteRecords()
	id := 1

	assert.Panics(t, func() {
		tests.UserService.Delete(id)
	})
}
