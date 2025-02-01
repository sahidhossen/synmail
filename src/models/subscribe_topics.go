package models

import (
	"time"

	"gorm.io/gorm"
)

type SubscribeTopics struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"` // Automatically updates on any change
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Automatically updates on any change
}

func (c *SubscribeTopics) TableName() string {
	return "subscribe_topics"
}
