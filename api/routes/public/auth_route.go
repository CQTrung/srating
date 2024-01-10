package public

import (
	"time"

	"srating/api/controllers"
	"srating/bootstrap"
	"srating/services"

	repositories "srating/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAuthRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		ur = repositories.NewUserRepository(db)
		au = services.NewAuthService(ur, timeout)
	)
	fc := controllers.AuthController{
		AuthService: au,
		Env:         env,
	}
	authGroup := group.Group("/auth")
	authGroup.POST("/register", fc.Register)
	authGroup.POST("/login", fc.Login)
	authGroup.POST("/refresh", fc.RefreshToken)
	// authGroup.POST("/change-password", fc.ChangePassword)
}
