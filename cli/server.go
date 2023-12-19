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
	asyn := app.AsynqClient
	// init
	// Initialize Redis and Timeout
	rd := app.Redis                                            // Obtain the Redis client instance from the 'app' object.
	timeout := time.Duration(env.RequestTimeout) * time.Second // Calculate the timeout duration.
	loc, err := time.LoadLocation(env.TimeZone)
	if err != nil {
		panic(err)
	}
	time.Local = loc

	// Configure Logger (zerolog)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix // Set the time field format to Unix time.
	zerolog.SetGlobalLevel(zerolog.InfoLevel)        // Set the global log level to Info.

	// Configure Swagger Documentation (docs)
	docs.SwaggerInfo.Title = "Tourist API"                                   // Set the title of the Swagger API documentation.
	docs.SwaggerInfo.Description = "This is a sample server Tourist server." // Set the API description.
	docs.SwaggerInfo.Version = "1.0"                                         // Set the API version.
	docs.SwaggerInfo.Host = env.BaseURL                                      // Set the API's base URL.
	docs.SwaggerInfo.BasePath = "/api/v1"                                    // Set the base path of API endpoints.
	docs.SwaggerInfo.Schemes = []string{"http", "https"}                     // Define supported schemes.
	// Initialize Router and Routes
	router := NewRouter(env)                         // Initialize the HTTP router.
	routes.Setup(env, timeout, router, db, rd, asyn) // Setup routes, middleware, and dependencies.

	// Configure and Start HTTP Server
	server := &http.Server{
		Addr:         env.ServerAddress,                             // Set the server address from configuration.
		Handler:      router,                                        // Set the HTTP router as the server handler.
		ReadTimeout:  time.Duration(env.ReadTimeout) * time.Second,  // Set the read timeout for incoming requests.
		WriteTimeout: time.Duration(env.WriteTimeout) * time.Second, // Set the write timeout for outgoing responses.
		IdleTimeout:  time.Duration(env.IdleTimeout) * time.Second,  // Set the idle timeout for persistent connections.
	}
	utils.LogInfo("Server started at " + env.ServerAddress) // Log the server start.

	// Signal Handling and Graceful Shutdown
	sigCh := make(chan os.Signal, 1)                      // Create a signal channel to capture termination/interruption signals.
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT) // Notify the channel for SIGTERM and SIGINT signals.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	<-sigCh // Block program execution until a termination/interruption signal is received.

	// Gracefully Shutdown Server
	ctx, cancel := context.WithTimeout(context.Background(), timeout) // Create a context with a timeout of 5 seconds.
	defer app.CloseConnection()
	defer cancel() // Defer the cancellation of the context.
	if err := server.Shutdown(ctx); err != nil {
		panic(err) // Panic if there's an error during graceful server shutdown.
	}

	utils.LogInfo("Server stopped") // Log the server stop.
}
