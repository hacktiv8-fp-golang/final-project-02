package main

import (
	"final-project-02/internal/database"
	"final-project-02/internal/router"
)

func main() {
	database.StartDB()
	router.StartServer()
}
