package services

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/topboyasante/mrkt-api/api/v1/dto"
	"github.com/topboyasante/mrkt-api/api/v1/models"
	"github.com/topboyasante/mrkt-api/api/v1/pkg/images"
	"github.com/topboyasante/mrkt-api/api/v1/repositories"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
)

type ListingService interface {
	CreateListing(r *dto.ListingServiceCreateRequest) (*models.Listing, error)
	GetListings() ([]models.Listing, error)
	GetListingByID(id string) (*models.Listing, error)
	GetListingsByUserID(id string) ([]models.Listing, error)
	GetFeaturedListings() ([]models.Listing, error)
	SearchListings(id string) ([]models.Listing, error)
	UpdateListing(listingID, userID string, data *dto.ListingServiceCreateRequest) (*models.Listing, error)
	DeleteListing(id string) error
}

type listingService struct {
	repo      repositories.ListingRepository
	validator *validator.Validate
}

func NewListingService(repo repositories.ListingRepository, validator *validator.Validate) ListingService {
	return &listingService{
		repo:      repo,
		validator: validator,
	}
}

func (s *listingService) CreateListing(r *dto.ListingServiceCreateRequest) (*models.Listing, error) {
	err := s.validator.Struct(r)
	if err != nil {
		return nil, err
	}

	file := r.Image
	if file.Size > 5242880 {
		return nil, errors.New("image size is above 5MB")
	}

	src, err := file.Open()
	if err != nil {
		utils.Logger().Error(err)
		return nil, err
	}
	defer src.Close()

	imgURL, imgPID, err := images.UploadToCloudinary(src, file.Filename)
	if err != nil {
		utils.Logger().Error(err)
		return nil, errors.New("unable to upload image")
	}

	l := &models.Listing{
		ID:            uuid.New().String(),
		UserID:        r.UserID,
		Title:         r.Title,
		Description:   r.Description,
		Price:         r.Price,
		Address:       r.Address,
		City:          r.City,
		Country:       r.Country,
		ImageURL:      imgURL,
		ImagePublicID: imgPID,
		IsActive:      true,
		IsFeatured:    false,
	}

	res, err := s.repo.Create(l)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *listingService) GetListings() ([]models.Listing, error) {
	listings, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return listings, nil
}

func (s *listingService) GetListingByID(id string) (*models.Listing, error) {
	err := s.validator.Var(id, "required")
	if err != nil {
		return nil, err
	}

	listing, err := s.repo.GetOneByIdentifier("id", id)
	if err != nil {
		return nil, err
	}

	return listing, nil
}

func (s *listingService) GetListingsByUserID(id string) ([]models.Listing, error) {
	err := s.validator.Var(id, "required")
	if err != nil {
		return nil, err
	}

	listings, err := s.repo.GetAllByIdentifier("user_id", id)
	if err != nil {
		return nil, err
	}

	return listings, nil
}

func (s *listingService) GetFeaturedListings() ([]models.Listing, error) {
	listings, err := s.repo.GetAllByIdentifier("is_featured", true)
	if err != nil {
		return nil, err
	}

	return listings, nil
}

func (s *listingService) SearchListings(query string) ([]models.Listing, error) {
	listings, err := s.repo.SearchListings(query)
	if err != nil {
		return nil, err
	}

	return listings, nil
}

func (s *listingService) UpdateListing(listingID, userID string, data *dto.ListingServiceCreateRequest) (*models.Listing, error) {
	l, err := s.repo.GetOneByIdentifier("id", listingID)
	if err != nil {
		return nil, err
	}

	if l.UserID != userID {
		return nil, errors.New("you are not the owner of this listing")
	}

	if data.Image != nil {
		file := data.Image
		if file.Size > 5242880 {
			return nil, errors.New("image size is above 5MB")
		}

		src, err := file.Open()
		if err != nil {
			utils.Logger().Error(err)
			return nil, err
		}
		defer src.Close()

		newImageURL, newImagePublicID, err := images.UpdateImageOnCloudinary(src, l.ImagePublicID, file.Filename)
		if err != nil {
			return nil, err
		}
		l.ImageURL = newImageURL
		l.ImagePublicID = newImagePublicID
	}

	if data.Title != "" {
		l.Title = data.Title
	}
	if data.Description != "" {
		l.Description = data.Description
	}
	if data.Price != 0 {
		l.Price = data.Price
	}
	if data.Address != "" {
		l.Address = data.Address

	}
	if data.City != "" {
		l.City = data.City

	}
	if data.Country != "" {
		l.Country = data.Country

	}

	res, err := s.repo.Update(l)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *listingService) DeleteListing(id string) error {
	l, err := s.repo.GetOneByIdentifier("id", id)
	if err != nil {
		return err
	}

	err = images.DeleteFromCloudinary(l.ImagePublicID)
	if err != nil {
		return err
	}

	err = s.repo.Delete(l.ID)
	if err != nil {
		return err
	}

	return nil
}
