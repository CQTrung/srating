package public

import (
	"time"

	"srating/api/controllers"
	"srating/bootstrap"
	"srating/services"

	repositories "srating/repositories"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewAuthRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB, rd *redis.Client, asyn *asynq.Client) {
	var (
		ur = repositories.NewUserRepository(db)
		au = services.NewAuthService(ur, asyn, timeout)
	)
	fc := controllers.AuthController{
		AuthService: au,
		Env:         env,
	}
	group.POST("/auth/register", fc.Register)
	group.POST("/auth/login", fc.Login)
	group.POST("/auth/refresh", fc.RefreshToken)
}
