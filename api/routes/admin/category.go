package admin

import (
	"time"

	"srating/api/controllers"
	"srating/bootstrap"
	repositories "srating/repositories"
	"srating/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewCategoryRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		fbr = repositories.NewCategoryRepository(db)
		fbs = services.NewCategoryService(fbr, timeout)
	)
	fc := controllers.CategoryController{
		CategoryService: fbs,
		Env:             env,
	}
	categoryGroup := group.Group("/categories")
	categoryGroup.GET("", fc.GetAllCategory)
	categoryGroup.POST("", fc.CreateCategory)
	categoryGroup.PUT("", fc.UpdateCategory)
	categoryGroup.DELETE("/:id", fc.DeleteCategory)
}
