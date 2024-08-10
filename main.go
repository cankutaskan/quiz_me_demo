package main

import (
	"quiz_me/api"
	"quiz_me/cmd"
	"quiz_me/db"
)

func main() {
	cmd.Execute()

	inMemoryDB := db.NewDBContext()
	db.Seed(inMemoryDB)

	apiServer := api.NewAPIServer(":8080", inMemoryDB)
	apiServer.Serve()
}
