package controllers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/topboyasante/mrkt-api/api/v1/dto"
	"github.com/topboyasante/mrkt-api/api/v1/services"
	"github.com/topboyasante/mrkt-api/api/v1/types"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
)

type ListingController struct {
	service services.ListingService
}

func NewListingController(service services.ListingService) *ListingController {
	return &ListingController{service: service}
}

func (c *ListingController) GetListings(ctx *gin.Context) {
	res, err := c.service.GetListings()
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusBadRequest, types.APIError{
			Message: "failed to get listings",
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: res,
	})
}

func (c *ListingController) GetFeaturedListings(ctx *gin.Context) {
	res, err := c.service.GetFeaturedListings()
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusBadRequest, types.APIError{
			Message: "failed to get featured listings",
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: res,
	})
}

func (c *ListingController) GetListingsByUserID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.service.GetListingsByUserID(id)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusBadRequest, types.APIError{
			Message: "failed to get listings",
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: res,
	})
}

func (c *ListingController) SearchListing(ctx *gin.Context) {
	id := ctx.Query("id")

	res, err := c.service.SearchListings(id)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusBadRequest, types.APIError{
			Message: "failed to get listings",
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: res,
	})
}

func (c *ListingController) GetListingByID(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := c.service.GetListingByID(id)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusBadRequest, types.APIError{
			Message: "failed to get listing",
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: res,
	})
}

func (c *ListingController) CreateListing(ctx *gin.Context) {
	var body dto.CreateListingRequest

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, types.APIError{Message: "unauthorized request"})
		return
	}

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusBadRequest, types.APIError{
			Message: "invalid file",
		})
		return
	}

	data := &dto.ListingServiceCreateRequest{
		UserID:      userID.(string),
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
		Address:     body.Address,
		City:        body.City,
		Country:     body.Country,
		Image:       file,
	}

	res, err := c.service.CreateListing(data)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "listing successfully created",
		Data:    res,
	})
}

func (c *ListingController) UpdateListing(ctx *gin.Context) {
	var body dto.CreateListingRequest

	id := ctx.Param("id")

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, types.APIError{Message: "unauthorized request"})
		return
	}

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	var file *multipart.FileHeader
	var err error

	// Attempt to get the image file, but do not return an error if it doesn't exist
	file, err = ctx.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusBadRequest, types.APIError{
			Message: "invalid file",
		})
		return
	}

	data := &dto.ListingServiceCreateRequest{
		UserID:      userID.(string),
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
		Address:     body.Address,
		City:        body.City,
		Country:     body.Country,
		Image:       file,
	}

	res, err := c.service.UpdateListing(id, userID.(string), data)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "listing updated",
		Data:    res,
	})
}

func (c *ListingController) DeleteListing(ctx *gin.Context) {
	id := ctx.Param("id")

	_, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, types.APIError{Message: "unauthorized request"})
		return
	}

	err := c.service.DeleteListing(id)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "listing successfully deleted",
	})
}
