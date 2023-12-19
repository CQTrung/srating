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

func NewDashboardRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB, rd *redis.Client, asyn *asynq.Client) {
	var (
		fbr = repositories.NewFeedbackRepository(db)
		fbs = services.NewFeedbackService(fbr, timeout)
		ur  = repositories.NewUserRepository(db)
		us  = services.NewUserService(ur, asyn, timeout)
		dbs = services.NewDashboardService(fbs, us, timeout)
	)
	fc := controllers.DashboardController{
		DashboardService: dbs,
		Env:              env,
	}
	group.GET("/dashboard", fc.Dashboard)
}
