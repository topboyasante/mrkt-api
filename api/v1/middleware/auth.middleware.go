package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/topboyasante/mrkt-api/api/v1/models"
	"github.com/topboyasante/mrkt-api/api/v1/types"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
	"github.com/topboyasante/mrkt-api/internal/config"
	"github.com/topboyasante/mrkt-api/internal/database"
	"gorm.io/gorm"
)

func RequireAuth(ctx *gin.Context) {
	tokenStr := ctx.GetHeader("Authorization")

	if tokenStr == "" {
		utils.Logger().Error("RequireAuth: Missing authorization header")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			utils.Logger().Error("RequireAuth: Unexpected signing method:", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.ENV.JWTKey), nil
	})

	if err != nil {
		utils.Logger().Error("RequireAuth: Error parsing token:", err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			utils.Logger().Error("RequireAuth: Token expired")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		err := database.DB.Where("id = ?", claims["sub"]).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				utils.Logger().Error("RequireAuth: error getting user by id: ", err)
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			utils.Logger().Error("RequireAuth: Error finding user:", err)
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("user_id", user.ID)
	} else {
		utils.Logger().Error("RequireAuth: Invalid token claims")
		ctx.JSON(http.StatusUnauthorized, types.APIError{Message: "Invalid token claims"})
	}
	ctx.Next()
}
