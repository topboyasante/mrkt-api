package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/topboyasante/mrkt-api/api/v1/dto"
	"github.com/topboyasante/mrkt-api/api/v1/services"
	"github.com/topboyasante/mrkt-api/api/v1/types"
	"github.com/topboyasante/mrkt-api/api/v1/utils"
)

type AuthController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) SignIn(ctx *gin.Context) {
	var body dto.SignInRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	res, err := c.service.SignIn(body)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "sign in successful",
		Data:    res,
	})
}

func (c *AuthController) SignUp(ctx *gin.Context) {
	var body dto.SignUpRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	res, err := c.service.SignUp(body)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "sign up successful",
		Data:    res,
	})
}

func (c *AuthController) ActivateAccount(ctx *gin.Context) {
	var body dto.ActivateAccountRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	err := c.service.ActivateAccount(body)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "account has been activated",
	})
}

func (c *AuthController) ForgotPassword(ctx *gin.Context) {
	var body dto.TokenRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	err := c.service.ForgotPassword(body)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "an email with a reset code has been sent to your email",
	})
}

func (c *AuthController) ResetPassword(ctx *gin.Context) {
	var body dto.ResetPasswordRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	err := c.service.ResetPassword(body)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "your password has been reset. please sign in",
	})
}

func (c *AuthController) RefreshAccessToken(ctx *gin.Context) {
	var body dto.RefreshAccessTokenRequest

	if err := ctx.Bind(&body); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	res, err := c.service.RefreshAccessToken(body)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Data: res,
	})
}

func (c *AuthController) ResendAuthToken(ctx *gin.Context) {
	var email string

	if err := ctx.Bind(&email); err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusOK, types.APIError{
			Message: "unable to parse request body",
		})
		return
	}

	err := c.service.ResendAuthToken(email)
	if err != nil {
		utils.Logger().Error(err)
		ctx.JSON(http.StatusUnauthorized, types.APIError{
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, types.APISuccess{
		Message: "auth token has been resent",
	})
}