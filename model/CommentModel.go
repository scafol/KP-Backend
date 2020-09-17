package model

import (
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/settings"
	"gorm.io/gorm/clause"
)

type CommentModel struct {
	db settings.DatabaseConfig
}

func (commentModel CommentModel) CreateComment(input *entity.Comment) (*entity.Comment, error) {
	err := commentModel.db.GetDatabaseConnection().Create(&input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (commentModel CommentModel) FindComments() ([]entity.Comment, error) {
	var comments []entity.Comment
	err := commentModel.db.GetDatabaseConnection().Preload(clause.Associations).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (commentModel CommentModel) FindComment(id string) (*entity.Comment, error) {
	var comment entity.Comment
	err := commentModel.db.GetDatabaseConnection().Find(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

func (commentModel CommentModel) UpdateComment(input *entity.Comment, id string) (*entity.Comment, error) {
	comment := &input
	err := commentModel.db.GetDatabaseConnection().Model(&input).Where("id = ?", id).Updates(comment).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (commentModel CommentModel) DeleteComment(id string) error {
	err := commentModel.db.GetDatabaseConnection().Delete(&entity.Comment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
