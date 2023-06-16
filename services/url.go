package services

import (
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/exception"
	"github.com/hisyamsk/url-shortener/helpers"
	"github.com/hisyamsk/url-shortener/models"
	"github.com/hisyamsk/url-shortener/repositories"
)

type UrlService interface {
	FindAll() []*models.UrlModel
	Find(field string, value any) *models.UrlModel
	Create(url *models.UrlModel) *models.UrlModel
	Update(url *models.UrlModel) *models.UrlModel
	Delete(id int)
}

type urlService struct {
	urlRepo  repositories.UrlRepository
	userRepo repositories.UserRepository
}

func NewUrlService(urlRepo repositories.UrlRepository, userRepo repositories.UserRepository) UrlService {
	return &urlService{urlRepo, userRepo}
}

func (service *urlService) FindAll() []*models.UrlModel {
	urls := service.urlRepo.FindAll()
	var urlsResponse []*models.UrlModel

	for _, val := range urls {
		urlsResponse = append(urlsResponse, helpers.UrlEntityToResponse(val))
	}

	return urlsResponse
}

func (service *urlService) Find(field string, value any) *models.UrlModel {
	url, err := service.urlRepo.Find(field, value)
	exception.RaiseIfNotFoundError(err)

	return helpers.UrlEntityToResponse(url)
}

func (service *urlService) Create(url *models.UrlModel) *models.UrlModel {
	// check if assigned user exists
	_, err := service.userRepo.Find("id", url.UserID)
	exception.RaiseIfNotFoundError(err)

	urlEntity := &entities.Url{Url: url.Url, Redirect: url.Redirect, UserID: url.UserID}
	// check if user send custom url
	if url.Url == "" {
		urlEntity.Url = helpers.GenerateRandomString()
	}

	// check if url already exists
	_, err = service.urlRepo.Find("url", url.Url)
	exception.RaiseIfDuplicateError(err, "url already exists")

	service.urlRepo.Create(urlEntity)

	return helpers.UrlEntityToResponse(urlEntity)
}

func (service *urlService) Update(url *models.UrlModel) *models.UrlModel {
	// check if url exists
	foundUrl, err := service.urlRepo.Find("id", url.ID)
	exception.RaiseIfNotFoundError(err)

	// check if assigned user exists
	_, err = service.userRepo.Find("id", url.UserID)
	exception.RaiseIfNotFoundError(err)

	// check if user updates the url
	if url.Url != "" && url.Url != foundUrl.Url {
		_, err = service.urlRepo.Find("url", url.Url)
		exception.RaiseIfDuplicateError(err, "url already exists")
	}

	urlEntity := &entities.Url{ID: url.ID, Url: url.Url, Redirect: url.Redirect, UserID: url.UserID}
	service.urlRepo.Update(urlEntity)

	return helpers.UrlEntityToResponse(urlEntity)
}

func (service *urlService) Delete(id int) {
	_, err := service.urlRepo.Find("id", id)
	exception.RaiseIfNotFoundError(err)

	service.urlRepo.Delete(id)
}
