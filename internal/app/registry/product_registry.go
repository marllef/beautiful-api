package registry

import (
	"marllef/beautiful-api/internal/app/api/controller"
	"marllef/beautiful-api/internal/app/services"
)

func (r *registry) NewProductController() controller.ProductController {
	productService := services.NewProductServices(r.container)
	return controller.NewProductController(productService)
}
