package public

import (
	"time"

	"srating/api/controllers"
	"srating/bootstrap"
	"srating/repositories"
	"srating/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewMediaRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		mr = repositories.NewMediaRepository(db)

		mu = services.NewMediaService(mr, timeout)
	)
	fc := controllers.MediaController{
		MediaService: mu,
		Env:          env,
	}
	group.POST("/media", fc.Upload)
}
