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

func UrlEntityToResponse(entity *entities.Url) *models.UrlModel {
	return &models.UrlModel{
		ID:       entity.ID,
		Url:      entity.Url,
		Redirect: entity.Redirect,
		UserID:   entity.UserID,
	}
}
