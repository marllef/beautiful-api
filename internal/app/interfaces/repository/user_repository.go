package repository

import (
	"marllef/beautiful-api/internal/app/models/dto"
	entities "marllef/beautiful-api/internal/app/models/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	Find(id int64) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	Create(data *dto.User) error
	Update(product *dto.User, id int64) error
	Delete(id int64) error
}

type userRepository struct {
	container *gorm.DB
	ProductRepository
}

func GetUserRepository(container *gorm.DB) (repo *userRepository) {
	db := container.Model(entities.Product{}).Session(&gorm.Session{})
	return &userRepository{
		container: db,
	}
}

func (r *userRepository) Create(user *dto.User) (err error) {
	result := r.container.Create(user)
	return result.Error
}

func (r *userRepository) Find(id int64) (user *entities.User, err error) {
	if result := r.container.Where("id = ?", id).First(&user); result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

func (r *userRepository) FindAll() (users []*entities.User, err error) {
	if result := r.container.Find(&users); result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *userRepository) Update(user *dto.User, id int64) (err error) {
	err = r.container.Where("id = ?", id).Select("*").Omit("id", "created_at").Updates(&entities.Product{
		Name: user.Name,
	}).Error

	return err
}

func (r *userRepository) Delete(id int64) (err error) {
	if result := r.container.Where("id = ?", id).Delete(&entities.Product{}); result.Error != nil {
		return result.Error
	}
	return nil
}
