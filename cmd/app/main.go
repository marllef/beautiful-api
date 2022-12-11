package main

import (
	"flag"

	"marllef/beautiful-api/configs"
	"marllef/beautiful-api/database"
	"marllef/beautiful-api/internal/app/registry"
	"marllef/beautiful-api/pkg/server"
)

// Main function
func main() {
	flag.Parse()

	err := configs.LoadEnvs()
	if err != nil {
		panic(err)
	}

	db, err := database.NewDB()
	if err != nil {
		return
	}

	reg := registry.NewRegistry(db)

	// Create new instance of server
	app := server.NewServer()

	// Add product routes on server
	product := reg.NewProductController()
	app.AddRoute("product", server.Route{
		Path:       "/",
		Controller: product.GetAll,
		Methods:    []string{"GET"},
	})

	app.AddRoute("products", server.Route{
		Path:       "/{id}",
		Controller: product.GetOne,
		Methods:    []string{"GET"},
	})

	// Start server
	app.Serve()
}
