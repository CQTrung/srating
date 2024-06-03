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
		fbcr = repositories.NewFeedbackCategoryRepository(db)
		ur   = repositories.NewUserRepository(db)
		fbcs = services.NewFeedbackCategoryService(fbcr, timeout)
		us   = services.NewUserService(ur, timeout)
		dbs  = services.NewDashboardService(fbcs, us, timeout)
	)
	fc := controllers.DashboardController{
		DashboardService: dbs,
		UserService:      us,
		Env:              env,
	}
	group.GET("/dashboard", fc.Dashboard)
}
