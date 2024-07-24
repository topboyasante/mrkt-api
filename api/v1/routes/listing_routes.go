package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/topboyasante/mrkt-api/api/v1/controllers"
	"github.com/topboyasante/mrkt-api/api/v1/middleware"
	"github.com/topboyasante/mrkt-api/api/v1/repositories"
	"github.com/topboyasante/mrkt-api/api/v1/services"
	"github.com/topboyasante/mrkt-api/internal/database"
)

func ListingRoutes(r *gin.RouterGroup, validator *validator.Validate) {
	listingRepository := repositories.NewListingRepository(database.DB)
	listingService := services.NewListingService(listingRepository, validator)
	listingController := controllers.NewListingController(listingService)

	listingRoutes := r.Group("/listings")
	listingRoutes.GET("", listingController.GetListings)
	listingRoutes.GET("/featured", listingController.GetFeaturedListings)
	listingRoutes.GET("/:id", listingController.GetListingByID)
	listingRoutes.GET("/user/:id", listingController.GetListingsByUserID)
	listingRoutes.GET("/search/:query", listingController.SearchListing)

	listingRoutes.Use(middleware.RequireAuth)
	listingRoutes.POST("/create", listingController.CreateListing)
	listingRoutes.PUT("/:id", listingController.UpdateListing)
	listingRoutes.DELETE("/:id", listingController.DeleteListing)
}
