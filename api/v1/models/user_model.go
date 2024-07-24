package models

import (
	"time"
)

type User struct {
	ID          string    `gorm:"column:id;primaryKey" json:"id,omitempty"`
	FirstName   string    `gorm:"column:first_name" json:"first_name,omitempty"`
	LastName    string    `gorm:"column:last_name" json:"last_name,omitempty"`
	Email       string    `gorm:"column:email" json:"email,omitempty"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number,omitempty"`
	Password    string    `gorm:"column:password" json:"password,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	IsActive    bool      `gorm:"column:is_active" json:"is_active,omitempty"`
	AuthToken   int       `gorm:"column:auth_token" json:"auth_token,omitempty"`
}

