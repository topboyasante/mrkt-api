package repositories

import (
	"errors"

	"github.com/topboyasante/mrkt-api/api/v1/models"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
	"gorm.io/gorm"
)

type ListingRepository interface {
	GetAll() ([]models.Listing, error)
	GetAllByIdentifier(identifier, id any) ([]models.Listing, error)
	GetOneByIdentifier(identifier, id any) (*models.Listing, error)
	SearchListings(query string) ([]models.Listing, error)
	Create(listing *models.Listing) (*models.Listing, error)
	Update(listing *models.Listing) (*models.Listing, error)
	Delete(id string) error
}

type listingRepository struct {
	db *gorm.DB
}

func NewListingRepository(db *gorm.DB) ListingRepository {
	return &listingRepository{
		db: db,
	}
}

func (r *listingRepository) GetAll() ([]models.Listing, error) {
	var listings []models.Listing

	res := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "first_name", "last_name", "email", "phone_number", "created_at")
	}).Find(&listings)
	if res.Error != nil {
		utils.Logger().Error("error getting listings: ", res.Error)
		return nil, res.Error
	}

	return listings, nil
}

func (r *listingRepository) GetAllByIdentifier(identifier, id any) ([]models.Listing, error) {
	var listings []models.Listing

	res := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "first_name", "last_name", "email", "phone_number", "created_at")
	}).Where(identifier, id).Find(&listings)

	if res.Error != nil {
		utils.Logger().Error("error getting listings by identifier: ", res.Error)
		return nil, res.Error
	}

	return listings, nil
}

func (r *listingRepository) GetOneByIdentifier(identifier, id any) (*models.Listing, error) {
	var listing models.Listing

	res := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "first_name", "last_name", "email", "phone_number", "created_at")
	}).Where(identifier, id).First(&listing)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("no listing exists with the provided identifier")
		}
		utils.Logger().Error("error getting listing by identifier: ", res.Error)
		return nil, res.Error
	}

	return &listing, nil
}

func (r *listingRepository) SearchListings(query string) ([]models.Listing, error) {
	var listings []models.Listing

	res := r.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "first_name", "last_name", "email", "phone_number", "created_at")
	}).Where("title ILIKE ?", "%"+query+"%").Find(&listings)
	if res.Error != nil {
		utils.Logger().Error("error retrieving listings: ", res.Error)
		return nil, res.Error
	}

	return listings, nil
}

func (r *listingRepository) Create(listing *models.Listing) (*models.Listing, error) {
	res := r.db.Create(&listing)

	if res.Error != nil {
		utils.Logger().Error("error creating listing: ", res.Error)
		return nil, res.Error
	}

	return listing, nil
}

func (r *listingRepository) Update(listing *models.Listing) (*models.Listing, error) {
	res := r.db.Save(&listing)

	if res.Error != nil {
		utils.Logger().Error("error updating listing: ", res.Error)
		return nil, res.Error
	}

	return listing, nil
}

func (r *listingRepository) Delete(id string) error {
	var listing models.Listing

	res := r.db.Where("id = ?", id).Delete(&listing)
	if res.Error != nil {
		utils.Logger().Error("error deleting listing: ", res.Error)
		return res.Error
	}

	return nil
}
