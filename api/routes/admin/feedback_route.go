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

func NewFeedbackRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		fbcr = repositories.NewFeedbackCategoryRepository(db)
		mr   = repositories.NewFeedbackRepository(db)
		fcs  = services.NewFeedbackCategoryService(fbcr, timeout)
		mu   = services.NewFeedbackService(mr, fcs, timeout)
	)
	fc := controllers.FeedbackController{
		FeedbackService: mu,
		Env:             env,
	}
	group.GET("/feedbacks", fc.GetAllFeedback)
	group.GET("/feedbacks/:id", fc.GetFeedbackDetail)
	group.POST("/feedbacks/level", fc.GetFeedbackByLevel)
}
