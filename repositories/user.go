package repositories

import (
	"fmt"

	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	FindAll() []*entities.User
	Find(field string, val any) (*entities.User, error)
	FindUrlsById(id int) []*entities.Url
	Create(user *entities.User)
	Update(user *entities.User)
	Delete(id int)
	DeleteUrlsById(id int)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct {
	DB *gorm.DB
}

func (repository *userRepository) FindAll() []*entities.User {
	var users []*entities.User
	err := repository.DB.Find(&users).Error
	helpers.PanicIfError(err)

	return users
}

func (repository *userRepository) Find(field string, val any) (*entities.User, error) {
	var user *entities.User
	query := fmt.Sprintf("%s = ?", field)
	err := repository.DB.Where(query, val).First(&user).Error

	return user, err
}

func (repository *userRepository) FindUrlsById(id int) []*entities.Url {
	var urls []*entities.Url
	err := repository.DB.Where("user_id = ?", id).Find(&urls).Error
	helpers.PanicIfError(err)

	return urls
}

func (repository *userRepository) Create(user *entities.User) {
	err := repository.DB.Create(&user).Error
	helpers.PanicIfError(err)
}

func (repository *userRepository) Update(user *entities.User) {
	err := repository.DB.Model(&user).Clauses(clause.Returning{}).Updates(&user).Error
	helpers.PanicIfError(err)
}

func (repository *userRepository) Delete(id int) {
	err := repository.DB.Delete(&entities.User{}, id).Error
	helpers.PanicIfError(err)
}

func (repository *userRepository) DeleteUrlsById(id int) {
	err := repository.DB.Where("user_id = ?", id).Delete(&entities.Url{}).Error
	helpers.PanicIfError(err)
}
