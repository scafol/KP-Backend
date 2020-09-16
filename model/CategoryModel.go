package model

import (
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/settings"
)

type CategoryModel struct {
	db settings.DatabaseConfig
}

func (CategoryModel CategoryModel) InsertCategory(Category entity.Category) (*entity.Category, error) {
	err := CategoryModel.db.GetDatabaseConnection().Create(&Category).Error
	if err != nil {
		return nil, err
	}
	return &Category, nil
}

func (CategoryModel CategoryModel) GetCategories() []entity.Category {
	var categories []entity.Category
	CategoryModel.db.GetDatabaseConnection().Find(&categories)
	return categories
}

func (CategoryModel CategoryModel) GetCategory(id string) entity.Category {
	var category entity.Category
	CategoryModel.db.GetDatabaseConnection().Find(&category, id)
	return category
}

func (CategoryModel CategoryModel) UpdateCategory(Category entity.Category, id string) (*entity.Category, error) {
	err := CategoryModel.db.GetDatabaseConnection().Model(&Category).Where("id = ?", id).Updates(Category).Error
	if err != nil {
		return nil, err
	}
	return &Category, nil
}

func (CategoryModel CategoryModel) DeleteCategory(id string) error {
	err := CategoryModel.db.GetDatabaseConnection().Delete(&entity.Category{}, id)
	if err != nil {
		return err.Error
	}
	return nil
}
