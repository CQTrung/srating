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

func NewLocationRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup, db *gorm.DB) {
	var (
		mr = repositories.NewLocationRepository(db)
		mu = services.NewLocationService(mr, timeout)
	)
	fc := controllers.LocationController{
		LocationService: mu,
		Env:               env,
	}
	locationGroup := group.Group("/locations")
	locationGroup.GET("", fc.GetAllLocation)
	locationGroup.POST("", fc.CreateLocation)
	locationGroup.PUT("", fc.UpdateLocation)
	locationGroup.DELETE("/:id", fc.DeleteLocation)
}
