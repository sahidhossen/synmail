package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Define Status Enum
type CampStatus string

const (
	DRAFT     CampStatus = "draft"
	SCHEDULED CampStatus = "scheduled"
	SENT      CampStatus = "sent"
	FAILED    CampStatus = "failed"
)

// Convert CampStatus to a format that GORM/Postgres understands
func (s *CampStatus) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid status type")
	}
	*s = CampStatus(str)
	return nil
}

// Convert Status to a database-friendly format
func (s CampStatus) Value() (driver.Value, error) {
	return string(s), nil
}

type Campaign struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	UserID      uint   `gorm:"not null"`
	TopicID     uint   `gorm:"not null"`
	TemplateID  uint   `gorm:"not null"`
	Name        string `binding:"required"`
	Subject     string `binding:"required"`
	Content     string
	Status      CampStatus `gorm:"type:varchar(20);default:'scheduled'"`
	ScheduledAt time.Time
	SendAt      time.Time
	CreatedAt   time.Time `gorm:"autoCreateTime"` // Automatically updates on any change
	UpdatedAt   time.Time `gorm:"autoUpdateTime"` // Automatically updates on any change
}

type UpdateCampaign struct {
	Name        string
	Subject     string
	Content     string
	Status      string
	ScheduledAt time.Time
	SendAt      time.Time
}

func (c *Campaign) TableName() string {
	return "campaigns"
}
