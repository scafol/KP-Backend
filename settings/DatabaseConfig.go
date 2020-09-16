package settings

import (
	"Golang-Echo-MVC-Pattern/entity"
	"log"

	"github.com/jinzhu/gorm"
)

type DatabaseConfig struct{}


// MySql Db Config
func (DatabaseConfig DatabaseConfig) GetDatabaseConfig() *gorm.DB {
	DB, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/discussion_crud?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	DB.Debug().AutoMigrate(
		entity.User{},
		entity.Discussion{},
		entity.Catagory{},
	)

	DB.Model(&entity.Discussion{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.Discussion{}).AddForeignKey("catagory_id", "catagory(id)", "CASCADE", "CASCADE")
	DB.Model(&entity.Catagory{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")

	return DB
}
