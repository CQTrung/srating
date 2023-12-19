package admin

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

func NewMediaRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB, rd *redis.Client, asyn *asynq.Client) {
	var (
		mr = repositories.NewMediaRepository(db)
		mu = services.NewMediaService(mr, timeout)
	)
	fc := controllers.MediaController{
		MediaService: mu,
		Env:          env,
	}
	group.GET("/media", fc.GetAll)
}
