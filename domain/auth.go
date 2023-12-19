package domain

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

type JwtCustomClaims struct {
	Role Role   `json:"role"`
	ID   string `json:"id"`
	jwt.RegisteredClaims
}
type JwtCustomRefreshClaims struct {
	ID   string `json:"id"`
	Role Role   `json:"role"`
	jwt.RegisteredClaims
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}
type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}
type AuthService interface {
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	Register(c context.Context, user *User) error
	Login(c context.Context, input *LoginRequest) (*User, error)
	ExtractIDFromToken(ctx context.Context, token, accessTokenSecret string) (uint, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
}
