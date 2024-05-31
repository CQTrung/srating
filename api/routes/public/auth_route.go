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
	group.POST("/auth/register", fc.Register)
	group.POST("/auth/login", fc.Login)
	group.POST("/auth/refresh", fc.RefreshToken)
}
