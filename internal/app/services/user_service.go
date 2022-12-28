package services

import (
	"marllef/beautiful-api/internal/app/interfaces/repository"
	entities "marllef/beautiful-api/internal/app/models/entity"

	"gorm.io/gorm"
)

type UserService interface {
	GetAllUsers() ([]*entities.User, error)
	GetOneUser(id int64) (*entities.User, error)
}

type userService struct {
	repository repository.UserRepository
	UserService
}

func NewUserServices(db *gorm.DB) *userService {
	repo := repository.GetUserRepository(db)
	return &userService{
		repository: repo,
	}
}

func (service *userService) GetAllUsers() ([]*entities.User, error) {
	users, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (service *userService) GetOneUser(id int64) (*entities.User, error) {
	products, err := service.repository.Find(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
