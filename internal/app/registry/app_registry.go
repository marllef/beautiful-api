package registry

import (
	"marllef/beautiful-api/internal/app/api/controller"
	"marllef/beautiful-api/internal/app/services"
	"marllef/beautiful-api/pkg/server"

	"gorm.io/gorm"
)

var applicationRegistry *appRegistry

type AppRegistry interface {
	NewProductController() controller.ProductController
	NewRouter() server.Router
	NewServer() server.Server
}

type appRegistry struct {
	container *gorm.DB
}

func NewAppRegistry(db *gorm.DB) *appRegistry {
	applicationRegistry = &appRegistry{
		container: db,
	}
	return applicationRegistry
}

func GetAppRegistry() *appRegistry {
	return applicationRegistry
}

func (r *appRegistry) NewProductController() controller.ProductController {
	productService := services.NewProductServices(r.container)
	return controller.NewProductController(productService)
}

func (r *appRegistry) NewServer() server.Server {
	return server.NewServer()
}

func (r *appRegistry) NewRouter() server.Router {
	rtr := server.NewRouter(server.NewServer(), r.NewProductController())
	return rtr
}
