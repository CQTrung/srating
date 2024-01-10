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
	userGroup := group.Group("/users")
	userGroup.POST("", fc.CreateUser)
	userGroup.GET("", fc.GetUserDetail)
	userGroup.PUT("/status", fc.ChangeStatus)
	userGroup.GET("/employees", fc.GetAllEmployee)
	userGroup.PUT("/employees", fc.UpdateEmployee)
	userGroup.POST("/change-password", fc.ChangePassword)
	userGroup.DELETE("/employees/:id", fc.DeleteEmployee)
	userGroup.POST("/reset-password", fc.ResetPassword)
}
