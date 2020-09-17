package entity

import (
	"gorm.io/gorm"
)

type Discussion struct {
	gorm.Model
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID uint   `json:"category_id"`
	UserID     uint   `json:"user_id"`
	User       User
	Comments   []Comment `gorm:"foreignKey:DiscussionID"`
}
