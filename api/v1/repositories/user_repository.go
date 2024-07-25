package repositories

import (
	"errors"

	"github.com/topboyasante/mrkt-api/api/v1/models"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByIdentifier(identifier, id string) (*models.User, error)
	GetPartialsByIdentifier(identifier, id string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetByIdentifier(identifier, id string) (*models.User, error) {
	var user models.User

	res := r.db.Where(identifier, id).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no user exists with the provided credentials")
		}
		utils.Logger().Error("error getting user by identifier: ", res.Error)
		return nil, res.Error
	}

	return &user, nil
}

func (r *userRepository) GetPartialsByIdentifier(identifier, id string) (*models.User, error) {
	var user models.User

	res := r.db.Select("id", "first_name", "last_name", "email", "phone_number", "created_at", "calling_code").Where(identifier, id).First(&user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no user exists with the provided credentials")
		}
		utils.Logger().Error("error getting user partials by identifier: ", res.Error)
		return nil, res.Error
	}

	return &user, nil
}

func (r *userRepository) Create(user *models.User) (*models.User, error) {
	res := r.db.Create(&user)

	if res.Error != nil {
		utils.Logger().Error("error creating user: ", res.Error)
		return nil, res.Error
	}

	return user, nil
}

func (r *userRepository) Update(user *models.User) (*models.User, error) {
	res := r.db.Save(&user)

	if res.Error != nil {
		utils.Logger().Error("error updating user: ", res.Error)
		return nil, res.Error
	}

	return user, nil
}

func (r *userRepository) Delete(id string) error {
	var user models.User

	res := r.db.Where("id = ?", id).Delete(&user)
	if res.Error != nil {
		utils.Logger().Error("error deleting user: ", res.Error)
		return res.Error
	}

	return nil
}
