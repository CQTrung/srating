package controllers

import (
	"srating/bootstrap"
	"srating/domain"
	"srating/x/rest"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService domain.AuthService
	Env         *bootstrap.Env
	*rest.JSONRender
}

// Register
// @Router /auth/register [post]
// @Tags auth
// @Summary Register user
// @Param request body domain.User true "request"
// @Success 200 {object} string
func (uc *AuthController) Register(c *gin.Context) {
	body := &domain.User{}
	err := c.ShouldBindJSON(&body)
	rest.AssertNil(err)
	err = uc.AuthService.Register(c, body)
	rest.AssertNil(err)
	uc.Success(c)
}

// Login
// @Router /auth/login [post]
// @Tags auth
// @Summary Login user
// @Param payload body domain.LoginRequest true "payload"
// @Success 200 {object} domain.LoginResponse
func (uc *AuthController) Login(c *gin.Context) {
	var body *domain.LoginRequest
	err := c.ShouldBindJSON(&body)
	rest.AssertNil(err)
	user, err := uc.AuthService.Login(c, body)
	rest.AssertNil(err)
	var (
		accessTokenSecret       = uc.Env.AccessTokenSecret
		refreshTokenSecret      = uc.Env.RefreshTokenSecret
		accessTokenExpiryHour   = uc.Env.AccessTokenExpiryHour
		refreshTokenExpiryHour  = uc.Env.RefreshTokenExpiryHour
		rememberTokenExpiryHour = uc.Env.RememberTokenExpiryHour
	)
	if body.IsRememberMe {
		accessTokenExpiryHour = rememberTokenExpiryHour
	}
	accessToken, err := uc.AuthService.CreateAccessToken(user, accessTokenSecret, accessTokenExpiryHour)
	rest.AssertNil(err)
	// Generate refresh token
	refreshToken, err := uc.AuthService.CreateRefreshToken(user, refreshTokenSecret, refreshTokenExpiryHour)
	rest.AssertNil(err)
	data := &domain.LoginResponse{
		ID:           user.ID,
		Username:     user.Username,
		Phone:        user.Phone,
		Email:        user.Email,
		ShortName:    user.ShortName,
		FullName:     user.FullName,
		Field:        user.Field,
		Avatar:       user.Avatar,
		Department:   user.Department,
		Role:         user.Role,
		Status:       domain.Status(user.Status),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	uc.SendData(c, data)
}

// RefreshToken
// @Router /auth/refresh [post]
// @Tags auth
// @Summary Refresh token
// @Param payload body domain.RefreshTokenRequest true "payload"
// @Success 200 {object} domain.RefreshTokenResponse
func (uc *AuthController) RefreshToken(c *gin.Context) {
	var input *domain.RefreshTokenRequest
	err := c.ShouldBindJSON(&input)
	rest.AssertNil(err)
	var (
		refreshTokenSecret    = uc.Env.RefreshTokenSecret
		accessTokenSecret     = uc.Env.AccessTokenSecret
		accessTokenExpiryHour = uc.Env.AccessTokenExpiryHour
	)
	userID, err := uc.AuthService.ExtractIDFromToken(c, input.RefreshToken, refreshTokenSecret)
	rest.AssertNil(err)
	user, err := uc.AuthService.GetUserByID(c, userID)
	rest.AssertNil(err)
	accessToken, err := uc.AuthService.CreateAccessToken(user, accessTokenSecret, accessTokenExpiryHour)
	rest.AssertNil(err)
	data := &domain.RefreshTokenResponse{
		AccessToken: accessToken,
	}
	uc.SendData(c, data)
}
