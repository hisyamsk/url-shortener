package repositories

import (
	"fmt"

	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UrlRepository interface {
	FindAll() []*entities.Url
	Find(field string, val any) (*entities.Url, error)
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

func (repository *urlRepository) Find(field string, val any) (*entities.Url, error) {
	var url *entities.Url
	query := fmt.Sprintf("%s = ?", field)
	err := repository.DB.Where(query, val).First(&url).Error

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
	err := repository.DB.Model(&url).Clauses(clause.Returning{}).Updates(&url).Error
	helpers.PanicIfError(err)
}

func (repository *urlRepository) Delete(id int) {
	var url *entities.Url
	repository.DB.Delete(&url, id)
}
