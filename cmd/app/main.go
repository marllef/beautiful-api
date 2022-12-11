package main

import (
	"flag"

	"marllef/beautiful-api/configs"
	"marllef/beautiful-api/database"
	"marllef/beautiful-api/internal/app/registry"
	logger "marllef/beautiful-api/pkg/mylogger"
	"marllef/beautiful-api/pkg/server"
)

var log = logger.Default()

// Main function
func main() {
	flag.Parse()

	err := configs.LoadEnvs()
	if err != nil {
		log.Errorf("Failed on read configs: %v", err)
		return
	}

	db, err := database.NewDB()
	if err != nil {
		log.Errorf("Failed on connect to database: %v", err)
		return
	}

	reg := registry.NewRegistry(db)

	// Create new instance of server
	app := server.NewServer()
	app.SetPrefix("/api")

	// Add product routes on server
	product := reg.NewProductController()
	app.AddRoute("product", server.Route{
		Path:       "/product",
		Controller: product.GetAll,
		Methods:    []string{"GET"},
	})

	app.AddRoute("products", server.Route{
		Path:       "/product/{id}",
		Controller: product.GetOne,
		Methods:    []string{"GET"},
	})

	// Start server
	err = app.Serve()
	if err != nil {
		log.Errorf("Failed to serve http server: %v", err)
	}
}
