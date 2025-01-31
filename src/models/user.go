package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserName  string    `json:"userName" binding:"required" gorm:"size:100;unique"`
	EmailID   string    `json:"emailId" binding:"required,email" gorm:"size:100;unique"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (c *User) TableName() string {
	return "users"
}
