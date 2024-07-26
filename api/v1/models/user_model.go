package models

import (
	"time"
)

type User struct {
	ID          string    `gorm:"column:id;primaryKey;not null" json:"id,omitempty"`
	FirstName   string    `gorm:"column:first_name;not null" json:"first_name,omitempty"`
	LastName    string    `gorm:"column:last_name;not null" json:"last_name,omitempty"`
	Email       string    `gorm:"column:email;unique;not null" json:"email,omitempty"`
	PhoneNumber string    `gorm:"column:phone_number;unique;not null" json:"phone_number,omitempty"`
	Password    string    `gorm:"column:password;not null" json:"password,omitempty"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
	IsActive    bool      `gorm:"column:is_active;not null" json:"is_active,omitempty"`
	AuthToken   int       `gorm:"column:auth_token" json:"auth_token,omitempty"`
}
