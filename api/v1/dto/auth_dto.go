package dto

import (
	"time"
)

type SignUpRequest struct {
	FirstName   string `json:"first_name" form:"first_name" validate:"required"`
	LastName    string `json:"last_name" form:"last_name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	CountryCode string `json:"country_code" form:"country_code" validate:"required,len=2"`
	Password    string `json:"password" form:"password" validate:"required,min=8"`
}

type UpdateUserRequest struct {
	FirstName   string `json:"first_name" form:"first_name" validate:"required"`
	LastName    string `json:"last_name" form:"last_name" validate:"required"`
	Email       string `json:"email" form:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" form:"phone_number" validate:"required"`
	CountryCode string `json:"country_code" form:"country_code" validate:"required,len=2"`
}

type SignInRequest struct {
	Email    string `json:"email" form:"email" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type TokenRequest struct {
	Email string `json:"email" form:"email" validate:"required,email"`
}

type SignUpResponse struct {
	ID          string    `json:"id" form:"id"`
	FirstName   string    `json:"first_name" form:"first_name"`
	LastName    string    `json:"last_name" form:"last_name"`
	Email       string    `json:"email" form:"email"`
	PhoneNumber string    `json:"phone_number" form:"phone_number"`
	CreatedAt   time.Time `json:"created_at" form:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" form:"updated_at,omitempty"`
	IsActive    bool      `json:"is_active" form:"is_active"`
	AuthToken   int       `json:"auth_token" form:"auth_token"`
}

type SignInResponse struct {
	FirstName    string `json:"first_name"`
	Email        string `json:"email"`
	ID           string `json:"id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type ActivateAccountRequest struct {
	Email     string `json:"email" validate:"required,email"`
	AuthToken int    `json:"auth_token" validate:"required"`
}

type ChangePasswordRequest struct {
	PrevPassword string `json:"prev_password" form:"prev_password" validate:"required,min=8"`
	NewPassword  string `json:"new_password" form:"new_password" validate:"required,min=8"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email" validate:"required,email"`
	AuthToken   int    `json:"auth_token" validate:"required"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}
