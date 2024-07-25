package dto

import "time"

type SignUpRequest struct {
	FirstName   string `json:"first_name" form:"first_name"`
	LastName    string `json:"last_name" form:"last_name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	CountryCode string `json:"country_code" form:"country_code"`
	CallingCode string `json:"calling_code" form:"calling_code"`
	Password    string `json:"password" form:"password"`
}
type UpdateUserRequest struct {
	FirstName   string `json:"first_name" form:"first_name"`
	LastName    string `json:"last_name" form:"last_name"`
	Email       string `json:"email" form:"email"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	CountryCode string `json:"country_code" form:"country_code"`
	CallingCode string `json:"calling_code" form:"calling_code"`
}

type SignInRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type TokenRequest struct {
	Email string `json:"email" form:"email"`
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
	RefreshToken string `json:"refresh_token"`
}

type RefreshAccessTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
}

type ActivateAccountRequest struct {
	Email     string `json:"email"`
	AuthToken int    `json:"auth_token"`
}

type ChangePasswordRequest struct {
	PrevPassword string `json:"prev_password" form:"prev_password"`
	NewPassword  string `json:"new_password" form:"new_password"`
}

type ResetPasswordRequest struct {
	Email       string `json:"email"`
	AuthToken   int    `json:"auth_token"`
	NewPassword string `json:"new_password"`
}
