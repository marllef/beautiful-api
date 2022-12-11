package main

import (
	"flag"

	"marllef/beautiful-api/configs"
	"marllef/beautiful-api/database"
	"marllef/beautiful-api/internal/app/registry"
	"marllef/beautiful-api/pkg/server"
)

// Init app function. Runs before the main function.
func init() {
	flag.Parse()

	err := configs.LoadEnvs()
	if err != nil {
		panic(err)
	}
}

// Main function
func main() {
	db, err := database.NewDB()
	if err != nil {
		return
	}
	app := registry.NewAppRegistry(db)
	router := app.NewRouter()
	httpServer := app.NewServer()
	productController := app.NewProductController()

	router.AddRoute("product", server.Route{
		Path:        "/",
		Middlewares: server.Middlewares{},
		Controller:  productController.GetAll,
		Methods:     []string{"GET"},
	})

	router.AddRoute("products", server.Route{
		Path:        "/{id}",
		Middlewares: server.Middlewares{},
		Controller:  productController.GetOne,
		Methods:     []string{"GET"},
	})

	httpServer.SetRoutes(router.GetRoutes())

	httpServer.Serve()
}
