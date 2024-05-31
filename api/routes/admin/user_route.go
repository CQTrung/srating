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

func NewUserRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		ur = repositories.NewUserRepository(db)
		uu = services.NewUserService(ur, timeout)
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
