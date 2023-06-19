package handlers

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/hisyamsk/url-shortener/tests"
	"github.com/stretchr/testify/assert"
)

func TestMainHandlerSuccess(t *testing.T) {
	tests.PopulateTables()
	defer tests.DeleteRecords()

	req := httptest.NewRequest(fiber.MethodGet, "/myurl3", nil)
	response, _ := tests.AppTest.Test(req, -1)

	assert.Equal(t, fiber.StatusFound, response.StatusCode)
}
