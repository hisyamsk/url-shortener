package repositories

import (
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
	"gorm.io/gorm"
)

type UrlRepository interface {
	FindAll() []*entities.Url
	FindById(id int) (*entities.Url, error)
	FindByUrl(value string) (*entities.Url, error)
	Create(url *entities.Url)
	Update(url *entities.Url)
	Delete(id int)
}

type urlRepository struct {
	DB *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{db}
}

func (repository *urlRepository) FindAll() []*entities.Url {
	var urls []*entities.Url
	err := repository.DB.Find(&urls).Error
	helpers.PanicIfError(err)

	return urls
}

func (repository *urlRepository) FindById(id int) (*entities.Url, error) {
	var url *entities.Url
	err := repository.DB.First(&url, id).Error

	return url, err
}

func (repository *urlRepository) FindByUrl(value string) (*entities.Url, error) {
	var url *entities.Url
	err := repository.DB.Where("url = ?", value).First(&url).Error

	return url, err
}

func (repository *urlRepository) Create(url *entities.Url) {
	err := repository.DB.Create(&url).Error
	helpers.PanicIfError(err)
}

func (repository *urlRepository) Update(url *entities.Url) {
	err := repository.DB.Save(&url).Error
	helpers.PanicIfError(err)
}

func (repository *urlRepository) Delete(id int) {
	var url *entities.Url
	repository.DB.Delete(&url, id)
}
