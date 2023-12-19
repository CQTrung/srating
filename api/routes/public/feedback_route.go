package public

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

func NewFeedbackRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB, rd *redis.Client, asyn *asynq.Client) {
	var (
		mr = repositories.NewFeedbackRepository(db)
		mu = services.NewFeedbackService(mr, timeout)
	)
	fc := controllers.FeedbackController{
		FeedbackService: mu,
		Env:             env,
	}
	group.POST("/feedbacks", fc.CreateFeedback)
}
