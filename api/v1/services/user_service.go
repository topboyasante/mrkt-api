package services

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/topboyasante/mrkt-api/api/v1/dto"
	"github.com/topboyasante/mrkt-api/api/v1/models"
	"github.com/topboyasante/mrkt-api/api/v1/pkg/auth"
	"github.com/topboyasante/mrkt-api/api/v1/pkg/email"
	"github.com/topboyasante/mrkt-api/api/v1/repositories"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
)

type UserService interface {
	GetUserByID(id string) (*models.User, error)
	UpdateUserDetails(id string, data *dto.UpdateUserRequest) (*models.User, error)
	ChangePassword(id string, data *dto.ChangePasswordRequest) (*models.User, error)
}

type userService struct {
	repo      repositories.UserRepository
	validator *validator.Validate
}

func NewUserService(repo repositories.UserRepository, validator *validator.Validate) UserService {
	return &userService{
		repo:      repo,
		validator: validator,
	}
}

func (s *userService) GetUserByID(id string) (*models.User, error) {
	user, err := s.repo.GetPartialsByIdentifier("id", id)
	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}
	return user, nil
}

func (s *userService) ChangePassword(id string, data *dto.ChangePasswordRequest) (*models.User, error) {
	user, err := s.repo.GetByIdentifier("id", id)
	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	err = auth.VerifyPassword(user.Password, data.PrevPassword)
	if err != nil {
		utils.Logger().Error(err)
		return nil, errors.New("the previous password is invalid")
	}

	hashedPassword, err := auth.HashPassword(data.NewPassword)
	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	user.Password = string(hashedPassword)
	usr, err := s.repo.Update(user)
	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	err = email.SendMailWithSMTP(
		email.EmailConfig,
		"Nana from MRKT",
		"Activate Your Account",
		"web/change-password.html",
		struct {
			Name      string
			AuthToken int
		}{Name: user.FirstName},
		[]string{user.Email},
	)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *userService) UpdateUserDetails(id string, data *dto.UpdateUserRequest) (*models.User, error) {
	user, err := s.repo.GetByIdentifier("id", id)
	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	if user.ID != id {
		return nil, errors.New("unauthorized operation. this is not your account")
	}

	if data.FirstName != "" {
		user.FirstName = data.FirstName
	}
	if data.LastName != "" {
		user.LastName = data.LastName
	}
	if data.PhoneNumber != "" {
		user.PhoneNumber = data.PhoneNumber
	}
	if data.Email != "" {
		user.Email = data.Email
	}
	if data.CountryCode != "" {
		user.CountryCode = data.CountryCode
	}
	if data.CallingCode != "" {
		user.CallingCode = data.CallingCode
	}
	
	res, err := s.repo.Update(user)
	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}

	return res, nil
}
