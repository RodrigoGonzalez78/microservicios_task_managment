package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Title       string `gorm:"not null; unique_index"  json:"title"`
	Description string `json:"description"`
	Done        bool   `gorm:"defaul:false" json:"done"`
	UserId      uint   `json:"user_id"`
}
