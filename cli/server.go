package cli

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"srating/api/middlewares"
	"srating/api/routes"
	"srating/bootstrap"
	"srating/docs"
	"srating/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(env *bootstrap.Env) *gin.Engine {
	gin.SetMode(env.GinMode)
	router := gin.New()
	router.Static("/statics", "./statics")
	router.Use(gin.Logger(), middlewares.AddHeader(), middlewares.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

func server(app *bootstrap.Application) *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			initServer(app)
		},
	}
}

func initServer(app *bootstrap.Application) {
	env := app.Env
	db := app.DB

	timeout := time.Duration(env.RequestTimeout) * time.Second
	loc, err := time.LoadLocation(env.TimeZone)
	if err != nil {
		panic(err)
	}
	time.Local = loc

	docs.SwaggerInfo.Title = "S-Rating API"
	docs.SwaggerInfo.Description = "This is a sample server S-Rating server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = env.BaseURL
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := NewRouter(env)
	routes.Setup(env, timeout, router, db)

	server := &http.Server{
		Addr:         env.ServerAddress,
		Handler:      router,
		ReadTimeout:  time.Duration(env.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(env.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(env.IdleTimeout) * time.Second,
	}
	utils.LogInfo("Server started at " + env.ServerAddress)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer app.CloseConnection()
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	utils.LogInfo("Server stopped")
}
