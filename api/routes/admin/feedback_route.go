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
		mr = repositories.NewFeedbackRepository(db)
		mu = services.NewFeedbackService(mr, timeout)
	)
	fc := controllers.FeedbackController{
		FeedbackService: mu,
		Env:             env,
	}
	// group.POST("/feedbacks", fc.CreateFeedback)
	group.GET("/feedbacks", fc.GetAllFeedback)
	group.GET("/feedbacks/:id", fc.GetFeedbackDetail)
	group.POST("/feedbacks/search", fc.SearchFeedback)
}
