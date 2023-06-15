package repositories_test

import (
	"testing"

	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
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

func TestUserRepoCreate(t *testing.T) {
	defer tests.DeleteRecords()
	result := &entities.User{Username: "foo", Password: "bar"}

	tests.UserRepo.Create(result)
	expected := &entities.User{ID: 1, Username: "foo", Password: "bar"}

	assert.Equal(t, expected, result)
}

func TestUserRepoFindById(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	result, err := tests.UserRepo.FindById(id)
	expected := tests.Users[0]

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
}

func TestUserRepoFindAll(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	result := tests.UserRepo.FindAll()
	expected := tests.Users

	assert.Equal(t, expected, result)
}

func TestUserRepoFindUrlsById(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	result := tests.UserRepo.FindUrlsById(id)
	var expected []*entities.Url
	for _, url := range tests.Urls {
		if url.UserID == uint(id) {
			expected = append(expected, url)
		}
	}

	assert.Equal(t, expected, result)
}

func TestUserRepoUpdate(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	result := tests.Users[0]
	result.Username = "updated_hisyam"

	tests.UserRepo.Update(result)
	expected := &entities.User{ID: 1, Username: "updated_hisyam", Password: "password1"}

	assert.Equal(t, expected, result)
}

func TestUserRepoDelete(t *testing.T) {
	tests.DB.Create(&entities.User{Username: "foo", Password: "baz"})
	defer tests.DeleteRecords()
	id := 1

	tests.UserRepo.Delete(id)
	_, err := tests.UserRepo.FindById(id)

	assert.NotNil(t, err)
}
