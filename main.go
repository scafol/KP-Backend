package main

import (
	"Golang-Echo-MVC-Pattern/routes"
	"Golang-Echo-MVC-Pattern/settings"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	// Database
	DB := settings.DatabaseConfig.GetDatabaseConfig(settings.DatabaseConfig{})
	defer DB.Close()

	// Starting server
	echo := routes.Routing.GetRoutes(routes.Routing{})
	err := echo.Start(":1337")
	if err != nil {
		log.Fatal(err)
	}

}
