package repositories_test

import (
	"testing"

	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/tests"
	"github.com/stretchr/testify/assert"
)

func TestUrlRepoCreate(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	result := &entities.Url{Url: "myurl", Redirect: "google.com", UserID: 1}

	tests.UrlRepo.Create(result)
	expected := &entities.Url{ID: 4, Url: "myurl", Redirect: "google.com", UserID: 1}

	assert.Equal(t, expected, result)
}

func TestUrlRepoFindById(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	result, err := tests.UrlRepo.FindById(id)
	expected := tests.Urls[0]

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestUrlRepoFindByUrl(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	url := "url1"

	result, err := tests.UrlRepo.FindByUrl(url)
	expected := tests.Urls[0]

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestUrlRepoFindAll(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	result := tests.UrlRepo.FindAll()
	expected := tests.Urls

	assert.Equal(t, expected, result)
}

func TestUrlRepoUpdate(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	result := tests.Urls[0]
	result.Redirect = "wikipedia.com"

	tests.UrlRepo.Update(result)
	expected := &entities.Url{ID: 1, Url: "url1", Redirect: "wikipedia.com", UserID: 1}

	assert.Equal(t, expected, result)
}

func TestUrlRepoDelete(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	tests.UrlRepo.Delete(id)
	_, err := tests.UrlRepo.FindById(id)

	assert.NotNil(t, err)
}
