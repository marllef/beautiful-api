package repository

import (
	"marllef/beautiful-api/internal/app/models/dto"
	entities "marllef/beautiful-api/internal/app/models/entity"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Find(id int64) (*entities.Product, error)
	FindAll() ([]*entities.Product, error)
	Create(data *dto.Product) error
	Update(product *dto.Product, id int64) error
	Delete(id int64) error
}

type productRepository struct {
	container *gorm.DB
	ProductRepository
}

func GetRepository(container *gorm.DB) (repo *productRepository) {
	db := container.Model(entities.Product{}).Session(&gorm.Session{})
	return &productRepository{
		container: db,
	}
}

func (r *productRepository) Create(data *dto.Product) (err error) {
	result := r.container.Create(data)
	return result.Error
}

func (r *productRepository) Find(id int64) (product *entities.Product, err error) {
	if result := r.container.Where("id = ?", id).First(&product); result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (r *productRepository) FindAll() (products []*entities.Product, err error) {
	if result := r.container.Find(&products); result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (r *productRepository) Update(product *dto.Product, id int64) (err error) {
	err = r.container.Where("id = ?", id).Select("*").Omit("id", "created_at").Updates(&entities.Product{
		Name: product.Name,
		Code: product.Code,
	}).Error

	return err
}

func (r *productRepository) Delete(id int64) (err error) {
	if result := r.container.Where("id = ?", id).Delete(&entities.Product{}); result.Error != nil {
		return result.Error
	}
	return nil
}
