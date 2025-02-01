package models

import (
	"time"

	"gorm.io/gorm"
)

type Trackers struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	CampaignID   uint   `gorm:"not null"`
	SubscriberID uint   `gorm:"not null"`
	Status       string `gorm:"default:sent"` //e.g., "Sent", "Delivered", "Bounced", "Opened"
	SendAt       time.Time
	OpenedAt     time.Time
	ClickedAt    time.Time
}
type TrackersUpdate struct {
	Status    string `gorm:"default:sent"` //e.g., "Sent", "Delivered", "Bounced", "Opened"
	SendAt    time.Time
	OpenedAt  time.Time
	ClickedAt time.Time
}

func (c *Trackers) TableName() string {
	return "trackers"
}
