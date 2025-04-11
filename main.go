package main

import (
	"log"

	"github.com/shivang-saxena/todo-app/config"
	"github.com/shivang-saxena/todo-app/routes"
)

func main() {
	// Load configurations
	config.LoadEnv()
	db := config.SetupDB()
	defer db.Close()

	// Setup router
	r := routes.SetupRouter(db)

	// Start server
	port := config.GetEnv("PORT", "8080")
	log.Printf("Server started on port %s", port)
	r.Run(":" + port)
}
