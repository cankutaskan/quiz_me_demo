package main

import (
	"quiz_me/api"
	"quiz_me/db"
)

func main() {
	// Initialize the in-memory database
	inMemoryDB := db.NewInMemoryDB()

	// Seed the database with initial data (if you have a seed function)
	db.Seed(inMemoryDB)

	// Create a new API server
	apiServer := api.NewAPIServer(":8080", inMemoryDB)

	// Start the server
	apiServer.Serve()
}
