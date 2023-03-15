package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID          uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Title       string         `json:"title" binding:"required"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
