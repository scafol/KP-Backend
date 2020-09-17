package entity

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ParentID     uint   `gorm:"index"`
	Content      string `json:"content"`
	DiscussionID uint   `json:"discussion_id"`
	UserID       uint   `json:"user_id"`
}
