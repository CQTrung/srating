package admin

import (
	"time"

	"srating/api/controllers"
	"srating/bootstrap"
	"srating/services"

	repositories "srating/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewDashboardRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		fbr = repositories.NewFeedbackRepository(db)
		fbs = services.NewFeedbackService(fbr, timeout)
		ur  = repositories.NewUserRepository(db)
		us  = services.NewUserService(ur, timeout)
		dbs = services.NewDashboardService(fbs, us, timeout)
	)
	fc := controllers.DashboardController{
		DashboardService: dbs,
		Env:              env,
	}
	group.GET("/dashboard", fc.Dashboard)
}
