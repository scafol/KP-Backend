package model

import (
	"Golang-Echo-MVC-Pattern/entity"
	"Golang-Echo-MVC-Pattern/settings"
	"fmt"
)


type DiscussionModel struct {
	db settings.DatabaseConfig
}

func (e *DiscussionModel) AddDiscussion(discussion *entity.Discussion) (*entity.Discussion, error) {
	err := e.db.GetDatabaseConfig().Save(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.AddDiscussion] error execute query %v \n", err)
		return nil, fmt.Errorf("failed add data discussion")
	}
	return discussion, nil
}

func (e *DiscussionModel) GetDiscussions()(*[]entity.Discussion, error) {
	var discussions []entity.Discussion
	err := e.db.GetDatabaseConfig().Find(&discussions).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.GetDiscussions] error execute query %v \n", err)
		return nil, fmt.Errorf("failed view all data discussion")
	}
	return &discussions, nil
}

func (e *DiscussionModel) GetDiscussionById(id int)(*entity.Discussion, error) {
	var discussion = entity.Discussion{}
	err := e.db.GetDatabaseConfig().Table("discussion").Where("id = ?", id).First(&discussion).Error
	if err != nil {
		fmt.Printf("[DiscussionModel.GetdiscussionById] error execute query %v \n", err)
		return nil, fmt.Errorf("id discussion is not exist")
	}
	return &discussion, nil
}