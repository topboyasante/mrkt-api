package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/topboyasante/mrkt-api/api/v1/dto"
	"github.com/topboyasante/mrkt-api/api/v1/models"
	"github.com/topboyasante/mrkt-api/api/v1/pkg/auth"
	"github.com/topboyasante/mrkt-api/api/v1/pkg/email"
	"github.com/topboyasante/mrkt-api/api/v1/repositories"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
	"github.com/topboyasante/mrkt-api/internal/config"
)

type AuthService interface {
	SignIn(r dto.SignInRequest) (*dto.SignInResponse, error)
	SignUp(r dto.SignUpRequest) (*dto.SignUpResponse, error)
	ActivateAccount(r dto.ActivateAccountRequest) error
	ForgotPassword(r dto.TokenRequest) error
	ResetPassword(r dto.ResetPasswordRequest) error
	RefreshAccessToken(r dto.RefreshAccessTokenRequest) (*dto.RefreshAccessTokenResponse, error)
}

type authService struct {
	repo      repositories.UserRepository
	validator *validator.Validate
}

func NewAuthService(repo repositories.UserRepository, validator *validator.Validate) AuthService {
	return &authService{
		repo:      repo,
		validator: validator,
	}
}

func (s *authService) SignIn(r dto.SignInRequest) (*dto.SignInResponse, error) {
	err := s.validator.Struct(r)
	if err != nil {
		return nil, err
	}

	usr, err := s.repo.GetByIdentifier("email", r.Email)
	if err != nil {
		return nil, err
	}

	if !usr.IsActive {
		return nil, errors.New("account has not been activated")
	}

	err = auth.VerifyPassword(usr.Password, r.Password)
	if err != nil {
		return nil, err
	}

	access_token, refresh_token, expiry, err := auth.CreateJWTTokens(usr.ID)
	if err != nil {
		return nil, err
	}

	usr.AuthToken = auth.GenerateAuthToken()
	_, err = s.repo.Update(usr)
	if err != nil {
		return nil, err
	}

	res := &dto.SignInResponse{
		FirstName:    usr.FirstName,
		Email:        usr.Email,
		ID:           usr.ID,
		AccessToken:  access_token,
		RefreshToken: refresh_token,
		ExpiresIn:    expiry,
	}
	return res, nil
}

func (s *authService) SignUp(r dto.SignUpRequest) (*dto.SignUpResponse, error) {
	err := s.validator.Struct(r)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := auth.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		ID:          uuid.NewString(),
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		PhoneNumber: r.PhoneNumber,
		Email:       r.Email,
		Password:    string(hashedPassword),
		AuthToken:   0,
	}

	usr, err := s.repo.Create(&user)
	if err != nil {
		return nil, err
	}

	authToken := auth.GenerateAuthToken()

	usr.AuthToken = authToken
	newUser, err := s.repo.Update(usr)
	if err != nil {
		return nil, err
	}

	err = email.SendMailWithSMTP(
		email.EmailConfig,
		"Nana from MRKT",
		"Activate Your Account",
		"web/activate-account.html",
		struct {
			Name      string
			AuthToken int
		}{Name: user.FirstName, AuthToken: authToken},
		[]string{r.Email},
	)
	if err != nil {
		return nil, err
	}

	res := &dto.SignUpResponse{
		ID:          newUser.ID,
		FirstName:   newUser.FirstName,
		LastName:    newUser.LastName,
		Email:       newUser.Email,
		PhoneNumber: newUser.PhoneNumber,
		IsActive:    newUser.IsActive,
		AuthToken:   newUser.AuthToken,
		CreatedAt:   newUser.CreatedAt,
		UpdatedAt:   newUser.UpdatedAt,
	}

	return res, nil
}

func (s *authService) ActivateAccount(r dto.ActivateAccountRequest) error {
	err := s.validator.Struct(r)
	if err != nil {
		return err
	}

	usr, err := s.repo.GetByIdentifier("email", r.Email)
	if err != nil {
		return err
	}

	if usr.IsActive {
		return errors.New("account already activated")
	}

	if usr.AuthToken != r.AuthToken {
		return errors.New("token is invalid")
	}

	usr.IsActive = true
	usr.AuthToken = 0

	_, err = s.repo.Update(usr)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ForgotPassword(forgotPassword dto.TokenRequest) error {
	err := s.validator.Struct(forgotPassword)
	if err != nil {
		return err
	}

	usr, err := s.repo.GetByIdentifier("email", forgotPassword.Email)
	if err != nil {
		return errors.New("user does not exist")
	}

	usr.AuthToken = auth.GenerateAuthToken()
	_, err = s.repo.Update(usr)
	if err != nil {
		return nil
	}

	err = email.SendMailWithSMTP(
		email.EmailConfig,
		"Nana from MRKT",
		"Forgot my Password",
		"web/activate-account.html",
		struct {
			Name      string
			AuthToken int
		}{Name: usr.FirstName, AuthToken: usr.AuthToken},
		[]string{forgotPassword.Email},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *authService) ResetPassword(resetPassword dto.ResetPasswordRequest) error {
	err := s.validator.Struct(resetPassword)
	if err != nil {
		return err
	}

	usr, err := s.repo.GetByIdentifier("email", resetPassword.Email)
	if err != nil {
		return errors.New("user does not exist")
	}

	if usr.AuthToken == resetPassword.AuthToken {
		return errors.New("token is invalid")
	}

	hashedPassword, err := auth.HashPassword(resetPassword.NewPassword)
	if err != nil {
		return err
	}

	usr.Password = string(hashedPassword)

	_, err = s.repo.Update(usr)
	if err != nil {
		return err
	}

	return err
}

func (s *authService) RefreshAccessToken(refreshAccessTokenRequest dto.RefreshAccessTokenRequest) (*dto.RefreshAccessTokenResponse, error) {
	err := s.validator.Struct(refreshAccessTokenRequest)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(refreshAccessTokenRequest.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ENV.JWTKey), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, err
	}

	userID := claims["sub"].(string)
	newAccessToken, newRefreshToken, newExpiry, err := auth.CreateJWTTokens(userID)
	if err != nil {
		utils.Logger().Errorf("failed to generate tokens: %v", err)
		return nil, err
	}

	res := &dto.RefreshAccessTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
		ExpiresIn:    newExpiry,
	}

	return res, nil
}
