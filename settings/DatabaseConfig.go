package settings

import (
	"fmt"

	"github.com/scafol/KP-Backend/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct{}

// GetDatabaseConnection - create database object
func (DatabaseConfig DatabaseConfig) GetDatabaseConnection() *gorm.DB {
	dsn := GoDotEnvVariable("CONN_STRING")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}
	database.Debug().AutoMigrate(
		&entity.User{},
		&entity.Category{},
	)
	return database
}
