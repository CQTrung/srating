package routes

import (
	"time"

	"srating/api/middlewares"
	"srating/api/routes/admin"
	"srating/api/routes/public"
	"srating/bootstrap"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, router *gin.Engine, db *gorm.DB) {
	// All Public API
	apiV1Router := router.Group("/api/v1")
	publicRouter := apiV1Router.Group("")

	// All Public API
	public.NewStatusRouter(env, timeout, publicRouter, db)
	public.NewAuthRouter(env, timeout, publicRouter, db)
	public.NewMediaRouter(env, timeout, publicRouter, db)
	public.NewFeedbackRouter(env, timeout, publicRouter, db)

	protectedRouter := apiV1Router.Group("")
	adminAPIRouter := protectedRouter.Group("")

	// All Admin API
	adminAPIRouter.Use(middlewares.JwtAuthMiddleware(env.AccessTokenSecret))
	admin.NewUserRouter(env, timeout, adminAPIRouter, db)
	admin.NewMediaRouter(env, timeout, adminAPIRouter, db)
	admin.NewFeedbackRouter(env, timeout, adminAPIRouter, db)
	admin.NewDashboardRouter(env, timeout, adminAPIRouter, db)
	admin.NewDepartmentRouter(env, timeout, adminAPIRouter, db)
	admin.NewCategoryRouter(env, timeout, adminAPIRouter, db)

	apiV2Router := router.Group("/api/v2")
	adminV2APIRouter := apiV2Router.Group("")
	public.NewFeedbackV2Router(env, timeout, adminV2APIRouter, db)
}
