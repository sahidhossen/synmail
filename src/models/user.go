package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	ADMIN Role = "admin"
	USER  Role = "user"
)

type User struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	UserName  string `binding:"required" gorm:"size:100;unique"`
	EmailID   string `binding:"required,email" gorm:"size:100;unique"`
	Password  string
	Role      string    `gorm:"default:'user'"`
	CreatedAt time.Time `gorm:"autoCreateTime"` // Automatically updates on any change
	UpdatedAt time.Time `gorm:"autoUpdateTime"` // Automatically updates on any change
}

type LoginRequest struct {
	UserName string
	EmailID  string
	Password string `binding:"required"`
}

func (c *User) TableName() string {
	return "users"
}
