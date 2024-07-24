package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/topboyasante/mrkt-api/api/v1/controllers"
	"github.com/topboyasante/mrkt-api/api/v1/repositories"
	"github.com/topboyasante/mrkt-api/api/v1/services"
	"github.com/topboyasante/mrkt-api/internal/database"
)

func AuthRoutes(r *gin.RouterGroup, validator *validator.Validate) {
	userRepository := repositories.NewUserRepository(database.DB)
	authService := services.NewAuthService(userRepository, validator)
	authController := controllers.NewAuthController(authService)
	
	authRoutes := r.Group("/auth")

	authRoutes.POST("/sign-in", authController.SignIn)
	authRoutes.POST("/sign-up", authController.SignUp)
	authRoutes.POST("/activate-account", authController.ActivateAccount)
	authRoutes.POST("/forgot-password", authController.ForgotPassword)
	authRoutes.POST("/reset-password", authController.ResetPassword)
	authRoutes.POST("/refresh-token", authController.RefreshAccessToken)
}