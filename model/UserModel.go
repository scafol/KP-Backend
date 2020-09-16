package model

import (
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/settings"
)

type UserModel struct {
	db settings.DatabaseConfig
}

func (UserModel UserModel) InsertUser(user *entity.User) (*entity.User, error) {
	err := UserModel.db.GetDatabaseConnection().Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (UserModel UserModel) GetUsers() []entity.User {
	var users []entity.User
	UserModel.db.GetDatabaseConnection().Find(&users)
	return users
}

func (UserModel UserModel) GetUser(id string) entity.User {
	var user entity.User
	UserModel.db.GetDatabaseConnection().Find(&user, id)
	return user
}
