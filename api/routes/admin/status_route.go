package admin

import (
	"time"

	"srating/bootstrap"

	api "github.com/appleboy/gin-status-api"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewStatusRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	group.GET("/status", api.GinHandler)
}
