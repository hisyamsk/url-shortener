package helpers

import (
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/models"
)

func UserEntityToResponse(entity *entities.User) *models.UserResponse {
	return &models.UserResponse{
		ID:       entity.ID,
		Username: entity.Username,
	}
}
