package domain

type LoginRequest struct {
	Username     string `json:"username" validate:"required"`
	Password     string `json:"password" validate:"required"`
	IsRememberMe bool   `json:"is_remember_me"`
}
type LoginResponse struct {
	ID           uint        `json:"id"`
	Username     string      `json:"username"`
	Phone        string      `json:"phone"`
	Email        string      `json:"email"`
	ShortName    string      `json:"short_name"`
	FullName     string      `json:"full_name"`
	Field        string      `json:"field"`
	Counter      string      `json:"counter"`
	Avatar       *Media      `json:"media"`
	Department   *Department `json:"department"`
	Location     *Location   `json:"location"`
	Role         Role        `json:"role"`
	Status       Status      `json:"status"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}

type LoginService interface {
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
