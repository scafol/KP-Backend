package model

import (
	"github.com/scafol/KP-Backend/entity"
	"github.com/scafol/KP-Backend/settings"
)

type ExampleModel struct {
	db settings.DatabaseConfig
}

// Get Example Post
func (ExampleModel ExampleModel) GetPosts() []entity.ExampleEntity {
	posts := []entity.ExampleEntity{
		{
			Title:       "NewsOne",
			Description: "NewsOneDescription",
		},
		{
			Title:       "NewsTwo",
			Description: "NewsTwoDescription",
		},
	}

	return posts
}
