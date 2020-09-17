package model

import (
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/settings"
	"gorm.io/gorm/clause"
)

type DiscussionModel struct {
	db settings.DatabaseConfig
}

func (discussionModel DiscussionModel) CreateDiscussion(input *entity.Discussion) (*entity.Discussion, error) {
	err := discussionModel.db.GetDatabaseConnection().Create(&input).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (discussionModel DiscussionModel) GetDiscussions() ([]entity.Discussion, error) {
	var discussions []entity.Discussion
	err := discussionModel.db.GetDatabaseConnection().Preload(clause.Associations).Find(&discussions).Error
	if err != nil {
		return nil, err
	}
	return discussions, nil
}

func (discussionModel DiscussionModel) GetDiscussion(id string) (*entity.Discussion, error) {
	var discussion entity.Discussion
	err := discussionModel.db.GetDatabaseConnection().Find(&discussion, id).Error
	if err != nil {
		return nil, err
	}
	return &discussion, nil
}

func (discussionModel DiscussionModel) UpdateDiscussion(input *entity.Discussion, id string) (*entity.Discussion, error) {
	discussion := &input
	err := discussionModel.db.GetDatabaseConnection().Model(&input).Where("id = ?", id).Updates(discussion).Error
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (discussionModel DiscussionModel) DeleteDiscussion(id string) error {
	err := discussionModel.db.GetDatabaseConnection().Delete(&entity.Discussion{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
