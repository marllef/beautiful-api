package main

import (
	"flag"

	"marllef/beautiful-api/configs"
	"marllef/beautiful-api/frameworks/database"
	"marllef/beautiful-api/frameworks/server"
	entities "marllef/beautiful-api/internal/app/models/entity"
	"marllef/beautiful-api/internal/app/registry"
	logger "marllef/beautiful-api/pkg/mylogger"
)

var log = logger.Default()

// Main function
func main() {
	flag.Parse()

	if err := configs.LoadEnvs(); err != nil {
		log.Errorf("Failed on read configs: %v", err)
		return
	}

	// Create new connection of database
	db, err := database.NewDatabase()
	if err != nil {
		log.Errorf("Failed on connect to database: %v", err)
		return
	}

	// Migrate database, add entities to migrate here
	if err = db.AutoMigrate(entities.Product{}, entities.User{}); err != nil {
		log.Errorf("Failed on migrate database: %v", err)
		return
	}

	// New registry instance
	reg := registry.NewRegistry(db)

	// Create new instance of server
	app := server.NewServer()
	app.SetPrefix("/api")

	// Registry routes
	routes := Routes(reg)
	
	// Set routes
	app.SetRoutes(routes)

	// Start server
	err = app.Serve()
	if err != nil {
		log.Errorf("Failed to serve http server: %v", err)
	}
}
