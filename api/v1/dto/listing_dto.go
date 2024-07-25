package dto

import "mime/multipart"

type ListingServiceCreateRequest struct {
	UserID      string                `json:"user_id" form:"user_id"`
	Title       string                `json:"title" form:"title"`
	Description string                `json:"description" form:"description"`
	Address     string                `json:"address" form:"address"`
	City        string                `json:"city" form:"city"`
	Country     string                `json:"country" form:"country"`
	Price       int                   `json:"price" form:"price"`
	Image       *multipart.FileHeader `json:"image" form:"image"`
}

type CreateListingRequest struct {
	Title       string `json:"title" form:"title" validate:"required"`
	Description string `json:"description" form:"description" validate:"required"`
	Address     string `json:"address" form:"address" validate:"required"`
	City        string `json:"city" form:"city" validate:"required"`
	Country     string `json:"country" form:"country" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required,gt=0"`
}
