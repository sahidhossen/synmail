package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Define Status Enum
type SubStatus string

const (
	SUBSCRIBED   SubStatus = "subscribed"
	UNSUBSCRIBED SubStatus = "unsubscribed"
	BOUNCED      SubStatus = "bounced"
)

// Convert SubStatus to a format that GORM/Postgres understands
func (s *SubStatus) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid status type")
	}
	*s = SubStatus(str)
	return nil
}

// Convert Status to a database-friendly format
func (s SubStatus) Value() (driver.Value, error) {
	return string(s), nil
}

type Subscriber struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null"`
	FirstName string
	LastName  string
	Email     string    `gorm:"unique;not null"`
	Status    SubStatus `gorm:"type:varchar(20);default:'subscribed'"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // Automatically updates on any change
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Automatically updates on any change
}

type UpdateSubscriber struct {
	FirstName string
	LastName  string
	Email     string    `gorm:"unique;not null"`
	Status    SubStatus `gorm:"type:varchar(20);default:'subscribed'"`
}

type SubscriberSchedulerData struct {
	ID    uint
	Email string
}

func (c *Subscriber) TableName() string {
	return "subscribers"
}
