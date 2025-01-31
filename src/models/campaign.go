package models

import "time"

type Campaign struct {
	ID          uint `gorm:"primaryKey"`
	UserID      uint `gorm:"not null"`
	Subject     string
	Content     string
	Status      string // Use an ENUM type in the database
	ScheduledAt time.Time
}

func (c *Campaign) TableName() string {
	return "campaigns"
}
