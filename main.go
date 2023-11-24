package main

import (
	"github.com/hacktiv8-fp-golang/final-project-02/internal/database"
	"github.com/hacktiv8-fp-golang/final-project-02/internal/router"
)

func main() {
	database.StartDB()
	router.StartServer()
}
