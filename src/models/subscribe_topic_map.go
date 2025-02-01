package models

import (
	"time"

	"gorm.io/gorm"
)

type SubscribeTopicMap struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	TopicID     uint      `gorm:"not null"`
	SubscribeID uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"` // Automatically updates on any change
}

func (c *SubscribeTopicMap) TableName() string {
	return "subscribe_topic_map"
}
