package models

import "time"

type Subscribe struct {
	ID        uint `gorm:"primaryKey"`
	Email     string
	CreatedAt time.Time
}

func (c *Subscribe) TableName() string {
	return "subscribes"
}
