package repositories

import (
	"github.com/hisyamsk/url-shortener/entities"
	"github.com/hisyamsk/url-shortener/helpers"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() []*entities.User
	FindById(id int) (*entities.User, error)
	Create(user *entities.User)
	Update(user *entities.User)
	Delete(id int)
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
func (repository *userRepository) FindById(id int) (*entities.User, error) {
	var user *entities.User
	err := repository.DB.First(&user, id).Error

	return user, err
}
func (repository *userRepository) Create(user *entities.User) {
	err := repository.DB.Create(&user).Error
	helpers.PanicIfError(err)
}
func (repository *userRepository) Update(user *entities.User) {
	err := repository.DB.Save(&user).Error
	helpers.PanicIfError(err)
}
func (repository *userRepository) Delete(id int) {
	err := repository.DB.Delete(&entities.User{}, id).Error
	helpers.PanicIfError(err)
}