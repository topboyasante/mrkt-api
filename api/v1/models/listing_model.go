package models

import (
	"time"
)

type Listing struct {
	ID            string    `gorm:"column:id;primaryKey" json:"id"`
	Title         string    `gorm:"column:title;not null" json:"title"`
	Price         int       `gorm:"column:price;not null" json:"price"`
	Description   string    `gorm:"column:description;not null" json:"description"`
	UserID        string    `gorm:"column:user_id" json:"user_id"`
	User          User      `json:"user"`
	Address       string    `gorm:"column:address;not null" json:"address"`
	City          string    `gorm:"column:city;not null" json:"city"`
	ImageURL      string    `gorm:"column:img_url;not null" json:"image_url"`
	ImagePublicID string    `gorm:"column:img_public_id;not null" json:"image_public_id"`
	Country       string    `gorm:"column:country;not null" json:"country"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsActive      bool      `gorm:"column:is_active" json:"is_active"`
	IsFeatured    bool      `gorm:"column:is_featured" json:"is_featured"`
}
