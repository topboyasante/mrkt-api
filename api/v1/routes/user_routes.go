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

func UserRoutes(r *gin.RouterGroup, validator *validator.Validate) {
	userRepository := repositories.NewUserRepository(database.DB)
	userService := services.NewUserService(userRepository, validator)
	userController := controllers.NewUserController(userService)

	listingRoutes := r.Group("/user")
	listingRoutes.GET("/:id", userController.GetUserByID)

	listingRoutes.Use(middleware.RequireAuth)
	listingRoutes.PUT("/update", userController.UpdateUser)
	listingRoutes.PUT("/change-password", userController.ChangePassword)

}
