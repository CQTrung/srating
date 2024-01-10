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

func NewDepartmentRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		mr = repositories.NewDepartmentRepository(db)
		mu = services.NewDepartmentService(mr, timeout)
	)
	fc := controllers.DepartmentController{
		DepartmentService: mu,
		Env:               env,
	}
	departmentGroup := group.Group("/departments")
	departmentGroup.GET("", fc.GetAllDepartment)
	departmentGroup.POST("", fc.CreateDepartment)
	departmentGroup.PUT("", fc.UpdateDepartment)
	departmentGroup.DELETE("/:id", fc.DeleteDepartment)
}
