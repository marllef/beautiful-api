package main

import (
	"marllef/beautiful-api/frameworks/server"
	"marllef/beautiful-api/internal/app/registry"
	logger "marllef/beautiful-api/pkg/mylogger"
)

func Routes(reg registry.Registry) server.Routes {
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

	bible, err := reg.NewBibleController()
	if err != nil {
		logger.Default().Errorf("Error on load bible: %v", err)
	} else {
		routes["bible"] = server.Route{
			Path:       "/book/{book}/chapter/{chapter}/verse/{verse}",
			Controller: bible.GetSingleVerse,
			Methods:    []string{"GET"},
		}
	}

	return routes
}
