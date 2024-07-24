package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/topboyasante/mrkt-api/api/v1/dto"
	"github.com/topboyasante/mrkt-api/api/v1/services"
	"github.com/topboyasante/mrkt-api/api/v1/types"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
)

type userController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *userController {
	return &userController{service: service}
}

func (c *userController) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var body dto.SignUpRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	usr, err := c.service.GetUserByID(id)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "failed to get user",
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: usr,
	})
}

func (c *userController) ChangePassword(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, types.APIError{Message: "unauthorized request"})
		return
	}

	var body dto.ChangePasswordRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	usr, err := c.service.ChangePassword(userID.(string), &body)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: usr,
	})
}

func (c *userController) UpdateUser(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, types.APIError{Message: "unauthorized request"})
		return
	}

	var body dto.UpdateUserRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	usr, err := c.service.UpdateUserDetails(userID.(string), &body)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "failed to update user",
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: usr,
	})
}
