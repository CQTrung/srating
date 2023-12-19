package public

import (
	"time"

	"srating/bootstrap"

	api "github.com/appleboy/gin-status-api"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewStatusRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB, rd *redis.Client, asyn *asynq.Client) {
	group.GET("/status", api.GinHandler)
}
