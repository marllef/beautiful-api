package services

import (
	"marllef/beautiful-api/internal/app/models/entity"
	"marllef/beautiful-api/internal/app/interfaces/repository"
	"gorm.io/gorm"
)

type ProductService interface {
	GetAllProducts() ([]*entities.Product, error)
	GetOneProduct(id int64) (*entities.Product, error)
}

type productService struct {
	repository repository.ProductRepository
	ProductService
}

func NewProductServices(db *gorm.DB) *productService {
	repo := repository.GetRepository(db)
	return &productService{
		repository: repo,
	}
}

func (service *productService) GetAllProducts() ([]*entities.Product, error) {
	products, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (service *productService) GetOneProduct(id int64) (*entities.Product, error) {
	products, err := service.repository.Find(id)
	if err != nil {
		return nil, err
	}
	return products, nil
}
