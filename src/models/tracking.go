package models

import "time"

type Trackers struct {
	ID           uint `gorm:"primaryKey"`
	CampaignID   uint `gorm:"not null"`
	SubscriberID uint `gorm:"not null"`
	OpenedAt     time.Time
}

func (c *Trackers) TableName() string {
	return "trackers"
}
