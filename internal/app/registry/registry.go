package registry

import (
	"marllef/beautiful-api/internal/app/api/controller"
	"gorm.io/gorm"
)

var applicationRegistry *registry

type Registry interface {
	NewProductController() controller.ProductController
	NewBibleController() (controller.BibleController, error)
}

type registry struct {
	container *gorm.DB
	Registry
}

func NewRegistry(db *gorm.DB) *registry {
	return &registry{
		container: db,
	}
}


