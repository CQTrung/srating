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

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB, rd *redis.Client, asyn *asynq.Client) {
	var (
		ur = repositories.NewUserRepository(db)
		uu = services.NewUserService(ur, asyn, timeout)
	)
	fc := controllers.UserController{
		UserService: uu,
		Env:         env,
	}
	group.POST("/users", fc.CreateUser)
	group.GET("/users", fc.GetUserDetail)
	group.PUT("/users/status", fc.ChangeStatus)
	group.GET("/users/employees", fc.GetAllEmployee)
	group.PUT("/users/employees", fc.UpdateEmployee)
	group.DELETE("/users/employees/:id", fc.DeleteEmployee)
}
