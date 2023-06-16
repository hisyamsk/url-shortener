package services

import (
	"testing"

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

func TestUserServiceCreate(t *testing.T) {
	user := &models.UserModel{Username: "foobar", Password: "barbaz"}
	defer tests.DeleteRecords()

	result := tests.UserService.Create(user)
	expected := &models.UserResponse{ID: 1, Username: "foobar"}

	assert.Equal(t, expected, result)
}

func TestUserServiceCreateFailed(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	user := &models.UserModel{Username: "hisyam", Password: "barbaz"}

	assert.Panics(t, func() {
		tests.UserService.Create(user)
	})
}

func TestUserServiceFindById(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	result := tests.UserService.Find("id", id)
	expected := helpers.UserEntityToResponse(tests.Users[0])

	assert.Equal(t, expected, result)
}

func TestUserServiceFindByIdFailed(t *testing.T) {
	defer tests.DeleteRecords()
	id := 1

	assert.Panics(t, func() {
		tests.UserService.Find("id", id)
	})
}

func TestUserServiceFindByUsername(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	username := "hisyam"

	result := tests.UserService.Find("username", username)
	expected := helpers.UserEntityToResponse(tests.Users[0])

	assert.Equal(t, expected, result)
}

func TestUserServiceFindByUsernameFailed(t *testing.T) {
	defer tests.DeleteRecords()
	username := "hisyam"

	assert.Panics(t, func() {
		tests.UserService.Find("username", username)
	})
}

func TestUserServiceFindAll(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	var expected []*models.UserResponse

	result := tests.UserService.FindAll()
	for _, val := range tests.Users {
		expected = append(expected, helpers.UserEntityToResponse(val))
	}

	assert.Equal(t, expected, result)
}

func TestUserServiceFindUrlsById(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	userId := 1
	var expected []*models.UrlModel

	result := tests.UserService.FindUrlsById(userId)
	for _, val := range tests.Urls {
		if val.UserID == uint(userId) {
			expected = append(expected, helpers.UrlEntityToResponse(val))
		}
	}

	assert.Equal(t, expected, result)
}

func TestUserServiceFindUrlsByIdFailed(t *testing.T) {
	defer tests.DeleteRecords()
	userId := 1

	assert.Panics(t, func() {
		tests.UserService.FindUrlsById(userId)
	})
}

func TestUserServiceUpdate(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	user := &models.UserModel{ID: 1, Username: "updated", Password: "hello"}

	result := tests.UserService.Update(user)
	expected := helpers.UserEntityToResponse(tests.Users[0])
	expected.Username = "updated"

	assert.Equal(t, expected, result)
}

func TestUserServiceUpdateFailedUsername(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	user := &models.UserModel{ID: 2, Username: "hisyam", Password: "hello"}

	assert.Panics(t, func() {
		tests.UserService.Update(user)
	})
}

func TestUserServiceUpdateFailedPassword(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	user := &models.UserModel{ID: 1, Username: "updated", Password: "password123"}

	assert.Panics(t, func() {
		tests.UserService.Update(user)
	})
}

func TestUserServiceDelete(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	tests.UserService.Delete(id)
	_, err := tests.UserRepo.Find("id", id)

	assert.NotNil(t, err)
}

func TestUserServiceDeleteFailed(t *testing.T) {
	defer tests.DeleteRecords()
	id := 1

	assert.Panics(t, func() {
		tests.UserService.Delete(id)
	})
}
