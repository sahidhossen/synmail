package models

import (
	"time"

	"gorm.io/gorm"
)

type Unsubscribers struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	CampaignID   uint `gorm:"not null"`
	SubscriberID uint `gorm:"not null"`
	Reason       string
	CreatedAt    time.Time `gorm:"autoCreateTime"` // Automatically updates on any change
}

type UnsubscribeUpdate struct {
	CampaignID   uint `json:"campaign_id"`
	SubscriberID uint `json:"subscriber_id"`
	Reason       string
}

func (c *Unsubscribers) TableName() string {
	return "unsubscribers"
}
