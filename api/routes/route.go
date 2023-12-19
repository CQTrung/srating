package routes

import (
	"time"

	"srating/api/middlewares"
	"srating/api/routes/admin"
	"srating/api/routes/public"
	"srating/bootstrap"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Setup(env *bootstrap.Env, timeout time.Duration, router *gin.Engine, db *gorm.DB, rd *redis.Client, asyn *asynq.Client) {
	// All Public API
	apiV1Router := router.Group("/api/v1")
	publicRouter := apiV1Router.Group("")

	// All Public API
	public.NewStatusRouter(env, timeout, publicRouter, db, rd, asyn)
	public.NewAuthRouter(env, timeout, publicRouter, db, rd, asyn)
	public.NewMediaRouter(env, timeout, publicRouter, db, rd, asyn)
	public.NewFeedbackRouter(env, timeout, publicRouter, db, rd, asyn)

	protectedRouter := apiV1Router.Group("")
	adminAPIRouter := protectedRouter.Group("/")

	// All Admin API
	adminAPIRouter.Use(middlewares.AdminAuthMiddleware(env.AccessTokenSecret))
	admin.NewUserRouter(env, timeout, adminAPIRouter, db, rd, asyn)
	admin.NewMediaRouter(env, timeout, adminAPIRouter, db, rd, asyn)
	admin.NewFeedbackRouter(env, timeout, adminAPIRouter, db, rd, asyn)
	admin.NewDashboardRouter(env, timeout, adminAPIRouter, db, rd, asyn)
}
