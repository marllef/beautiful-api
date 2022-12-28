package main

import (
	"marllef/beautiful-api/frameworks/server"
	"marllef/beautiful-api/internal/app/registry"
)

func Routes(reg registry.Registry) (server.Routes) {
	routes := make(server.Routes)

	// Set product routes
	product := reg.NewProductController()
	routes["product"] = server.Route{
		Path:       "/product",
		Controller: product.GetAll,
		Methods:    []string{"GET"},
	}
	
	routes["products"] = server.Route{
		Path:       "/product/{id}",
		Controller: product.GetOne,
		Methods:    []string{"GET"},
	}

	return routes
}