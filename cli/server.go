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
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"gorm.io/gorm"
)

func NewRouter(env *bootstrap.Env) *gin.Engine {
	gin.SetMode(env.GinMode)
	router := gin.New()
	router.Static("/assets", "./assets")
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
	// asyn := app.AsynqClient

	// Initialize Redis and Timeout
	initializeEnvironment(env)

	// Configure Logger (zerolog)
	configureLogger()

	// Configure Swagger Documentation (docs)
	configureSwagger(env)

	// Initialize Router and Routes
	router := initializeRouter(env, db)

	// Configure and Start HTTP Server
	server := configureHTTPServer(env, router)

	startServer(env, server)
}

func initializeEnvironment(env *bootstrap.Env) {
	loc, err := time.LoadLocation(env.TimeZone)
	if err != nil {
		panic(err)
	}
	time.Local = loc
}

func configureLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func configureSwagger(env *bootstrap.Env) {
	docs.SwaggerInfo.Title = "S-Rating API"
	docs.SwaggerInfo.Description = "This is a sample server S-Rating server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = env.BaseURL
	docs.SwaggerInfo.BasePath = "/api/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func initializeRouter(env *bootstrap.Env, db *gorm.DB) http.Handler {
	router := NewRouter(env)
	routes.Setup(env, time.Duration(env.RequestTimeout)*time.Second, router, db)
	return router
}

func configureHTTPServer(env *bootstrap.Env, router http.Handler) *http.Server {
	return &http.Server{
		Addr:         env.ServerAddress,
		Handler:      router,
		ReadTimeout:  time.Duration(env.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(env.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(env.IdleTimeout) * time.Second,
	}
}

func startServer(env *bootstrap.Env, server *http.Server) {
	utils.LogInfo("Server started at " + env.ServerAddress)

	// Signal Handling and Graceful Shutdown
	sigCh := make(chan os.Signal, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	<-sigCh
	utils.LogInfo("Shutting down")
	// Gracefully Shutdown Server
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(env.RequestTimeout)*time.Second)
	// defer app.CloseConnection()
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	utils.LogInfo("Shutdown gracefully")
}
