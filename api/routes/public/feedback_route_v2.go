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

func NewFeedbackV2Router(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		fcr = repositories.NewFeedbackCategoryRepository(db)
		mr  = repositories.NewFeedbackRepository(db)
		fcs = services.NewFeedbackCategoryService(fcr, timeout)
		mu  = services.NewFeedbackService(mr, fcs, timeout)
	)
	fc := controllers.FeedbackController{
		FeedbackService: mu,
		Env:             env,
	}
	group.POST("/feedbacks", fc.CreateFeedbackV2)
}
