package domain

type LoginRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	IsRememberMe bool   `json:"is_remember_me"`
}
type LoginResponse struct {
	ID           uint   `json:"id"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	ShortName    string `json:"name"`
	Username     string `json:"username"`
	Role         Role   `json:"role"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginService interface {
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
