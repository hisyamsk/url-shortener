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

func TestUserServiceFindById(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	result := tests.UserService.Find("id", id)
	expected := helpers.UserEntityToResponse(tests.Users[0])

	assert.Equal(t, expected, result)
}

func TestUserServiceFindByUsername(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	username := "hisyam"

	result := tests.UserService.Find("username", username)
	expected := helpers.UserEntityToResponse(tests.Users[0])

	assert.Equal(t, expected, result)
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

func TestUserServiceUpdate(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	user := &models.UserModel{ID: 1, Username: "updated", Password: "hello"}

	result := tests.UserService.Update(user)
	expected := helpers.UserEntityToResponse(tests.Users[0])
	expected.Username = "updated"

	assert.Equal(t, expected, result)
}

func TestUserServiceDelete(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()
	id := 1

	tests.UserService.Delete(id)
	_, err := tests.UserRepo.Find("id", id)

	assert.NotNil(t, err)
}
