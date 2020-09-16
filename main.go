package main

import (
	"fmt"

	"github.com/scafol/KP-Backend/routes"
	"github.com/scafol/KP-Backend/settings"
)

// Starting server
func main() {
	// running automigrate
	db, err := settings.DatabaseConfig{}.GetDatabaseConnection().DB()
	if err != nil {
		fmt.Printf("error when connecting database : %s", err.Error())
	}
	defer db.Close()

	echo := routes.Routing.GetRoutes(routes.Routing{})
	_ = echo.Start(settings.GoDotEnvVariable("LISTEN_PORT"))
}
