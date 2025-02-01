package models

import (
	"time"

	"gorm.io/gorm"
)

type Template struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	Name      string
	Subject   string `json:"subject" binding:"required"`
	Content   string
	Status    string    `gorm:"default:publish"` // draft / publish
	CreatedAt time.Time `gorm:"autoCreateTime"`  // Automatically updates on any change
	UpdatedAt time.Time `gorm:"autoUpdateTime"`  // Automatically updates on any change
}

type UpdateTemplate struct {
	Name    string
	Subject string
	Content string
	Status  string
}

func (c *Template) TableName() string {
	return "templates"
}
